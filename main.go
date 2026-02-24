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

func main() {
	ui := RunTUI()
	fmt.Printf("\nStarting search...\n")
	fmt.Printf("Folder: %s\n", ui.folder)
	fmt.Printf("Column: %s\n", ui.column)
	fmt.Printf("Skills: %s\n", ui.skills)
	fmt.Printf("Mode: %s\n", ui.mode)
	if ui.mode == "semantic" {
		fmt.Printf("Threshold: %s\n", ui.threshold)
	}
	fmt.Println()

	var model *Model
	querySkills := parseSkills(ui.skills)
	fmt.Printf("Parsed query skills: %v\n", querySkills)
	var queryVec []float64
	var threshold float64

	if ui.mode == "semantic" {
		fmt.Println("Preparing semantic search...")
		threshold, _ = strconv.ParseFloat(ui.threshold, 64)
		fmt.Printf("Using threshold: %.2f\n", threshold)

		err := EnsureModelExists()
		if err != nil {
			log.Fatal("Failed to prepare embedding model:", err)
		}

		fmt.Println("Loading embedding model...")
		model, err = LoadModel("model.vec")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Computing query embedding...")
		var vecs [][]float64
		for _, s := range querySkills {
			v := model.Embed(s)
			if v != nil {
				vecs = append(vecs, v)
				fmt.Printf("Found embedding for '%s' (dim: %d)\n", s, len(v))
			} else {
				fmt.Printf("No embedding found for '%s'\n", s)
			}
		}
		queryVec = averageVec(vecs)
		fmt.Printf("Query vector computed (dim: %d)\n", len(queryVec))

		// Test similarity with identical string
		if len(querySkills) > 0 {
			testVec := model.Embed(querySkills[0])
			if testVec != nil && len(queryVec) > 0 {
				testScore := cosine(queryVec, testVec)
				fmt.Printf("Self-similarity test for '%s': %.6f\n", querySkills[0], testScore)
			}
		}
	}

	fmt.Println("Scanning for Excel files...")
	files, _ := filepath.Glob(filepath.Join(ui.folder, "*.xlsx"))
	fmt.Printf("Found %d Excel files to process\n", len(files))

	var wg sync.WaitGroup
	var mu sync.Mutex
	var results [][]string
	processedFiles := 0

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
				fmt.Printf("Column '%s' not found in file %s (available columns: %v)\n", ui.column, filepath.Base(file), rows[0])
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
					vec := model.Embed(cell)
					if vec == nil {
						continue
					}
					score := cosine(queryVec, vec)

					// Debug high-scoring matches
					if score > 0.8 {
						fmt.Printf("High score match: '%s' -> score: %.3f\n", cell, score)
					}

					if score >= threshold {
						mu.Lock()
						results = append(results, row)
						localMatches++
						mu.Unlock()
					}
				}
			}

			if localMatches > 0 {
				fmt.Printf("File %s: found %d matches\n", filepath.Base(file), localMatches)
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
