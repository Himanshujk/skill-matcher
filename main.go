// Package main implements a skill matching application that searches for skills across Excel files
// using either exact text matching or semantic similarity with word embeddings.
// The application supports concurrent processing of multiple Excel files and exports results
// to timestamped Excel files.
// Package main implements a skill matching application that searches for skills across Excel files
// using either exact text matching or semantic similarity with word embeddings.
// The application supports concurrent processing of multiple Excel files and exports results
// to timestamped Excel files.
package main

import (
	"fmt"
	"log"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/xuri/excelize/v2"
)

// main is the entry point of the skill matcher application.
// It orchestrates the entire workflow:
// 1. Collects user input via TUI
// 2. Prepares semantic search models if needed
// 3. Processes Excel files concurrently
// 4. Exports matching results to a new Excel file
func main() {
	// Run interactive terminal UI to collect search parameters
	ui := RunTUI()

	// Clear terminal and reset cursor position for clean output
	// \033[2J clears entire screen, \033[H moves cursor to top-left
	fmt.Print("\033[2J\033[H")

	// Display search configuration for user confirmation
	fmt.Printf("Starting search...\n")
	fmt.Printf("Folder: %s\n", ui.folder)
	fmt.Printf("Column: %s\n", ui.column)
	fmt.Printf("Skills: %s\n", ui.skills)
	fmt.Printf("Mode: %s\n", ui.mode)
	if ui.mode == "exact + semantic" {
		fmt.Printf("Threshold: %s\n", ui.threshold)
	}
	fmt.Println()

	// Initialize variables for search processing
	var model *Model
	querySkills := parseSkills(ui.skills)
	var queryVec []float64 // Vector representation of query skills
	var threshold float64  // Similarity threshold for semantic matching

	if ui.mode == "exact + semantic" {
		fmt.Println("Preparing exact + semantic search...")
		threshold, _ = strconv.ParseFloat(ui.threshold, 64)
		fmt.Printf("Using threshold: %.2f\n", threshold)

		// Download model if it doesn't exist locally
		err := EnsureModelExists()
		if err != nil {
			log.Fatal("Failed to prepare embedding model:", err)
		}

		// Load pre-trained GloVe word embeddings
		fmt.Println("Loading embedding model...")
		model, err = LoadModel("model.vec")
		if err != nil {
			log.Fatal(err)
		}

		// Convert query skills to vector representations and average them
		// This creates a single query vector that represents all input skills
		fmt.Println("Computing query embedding...")
		// Collect vectors for all query skills that have embeddings
		var vecs [][]float64
		for _, s := range querySkills {
			v := model.Embed(s)
			if v != nil {
				vecs = append(vecs, v)
			}
		}
		// Average all skill vectors to create a single query vector
		queryVec = averageVec(vecs)
		fmt.Printf("Query vector computed (dim: %d)\n", len(queryVec))
	}

	// Find all Excel files in the specified directory
	fmt.Println("Processing Excel files...")
	files, _ := filepath.Glob(filepath.Join(ui.folder, "*.xlsx"))
	fmt.Printf("Found %d Excel files to process\n", len(files))

	// Setup concurrent processing for multiple Excel files
	var wg sync.WaitGroup  // Wait group for goroutine synchronization
	var mu sync.Mutex      // Mutex for thread-safe access to shared data
	var results [][]string // Collection of all matching rows
	processedFiles := 0    // Counter for progress tracking

	// Process each Excel file in a separate goroutine for better performance
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			defer func() {
				mu.Lock()
				processedFiles++
				fmt.Printf("\rProcessed %d/%d files", processedFiles, len(files))
				mu.Unlock()
			}()

			f, err := excelize.OpenFile(file)
			if err != nil {
				return
			}

			sheet := f.GetSheetName(0)
			rows, _ := f.GetRows(sheet)

			if len(rows) == 0 {
				return
			}

			colIndex := -1
			for i, h := range rows[0] {
				if h == ui.column {
					colIndex = i
					break
				}
			}

			if colIndex == -1 {
				return
			}

			localMatches := 0
			for _, row := range rows[1:] {
				if colIndex >= len(row) {
					continue
				}

				cell := strings.TrimSpace(row[colIndex])

				if ui.mode == "exact" {
					cellSkills := parseSkills(cell)
					// Check if any query skill matches any cell skill
					found := false
					for _, querySkill := range querySkills {
						if slices.Contains(cellSkills, querySkill) {
							found = true
							break
						}
					}
					if found {
						mu.Lock()
						results = append(results, row)
						localMatches++
						mu.Unlock()
					}
				} else {
					// Exact + semantic mode: try both approaches
					found := false

					// First try exact matching
					cellSkills := parseSkills(cell)
					for _, querySkill := range querySkills {
						if slices.Contains(cellSkills, querySkill) {
							found = true
							break
						}
					}

					// If not found by exact match, try semantic matching
					if !found {
						vec := model.Embed(cell)
						if vec != nil {
							score := cosine(queryVec, vec)
							if score >= threshold {
								found = true
							}
						}
					}

					if found {
						mu.Lock()
						results = append(results, row)
						localMatches++
						mu.Unlock()
					}
				}
			}

			if localMatches > 0 {
				mu.Lock()
				fmt.Printf("\nFile %s: found %d matches\n", filepath.Base(file), localMatches)
				mu.Unlock()
			}
		}(file)
	}

	wg.Wait()
	fmt.Printf("\nProcessing complete!\n")
	fmt.Printf("Found %d matching results\n", len(results))

	if len(results) == 0 {
		fmt.Println("No matches found. Try adjusting your search criteria.")
		return
	}

	fmt.Println("Writing results to Excel file...")

	// Generate timestamp-based filename
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("results_%s.xlsx", timestamp)

	out := excelize.NewFile()
	sheet := out.GetSheetName(0)

	for i, row := range results {
		for j, val := range row {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+1)
			err := out.SetCellValue(sheet, cell, val)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err := out.SaveAs(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n✅ Search Complete! Output written to %s (%d matches found)\n", filename, len(results))
}
