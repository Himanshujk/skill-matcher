// Package main handles downloading and extracting pre-trained word embedding models.
// Specifically designed for Stanford's GloVe (Global Vectors) word embeddings.
// The download process includes progress tracking via terminal UI.
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

// Model download configuration
const modelURL = "https://nlp.stanford.edu/data/wordvecs/glove.2024.wikigiga.100d.zip" // Stanford GloVe download URL
const zipFile = "glove.2024.wikigiga.100d.zip"                                         // Local zip file name

// targetFile will be updated when we discover the actual filename inside the zip
var targetFile = "glove.2024.wikigiga.100d.txt" // Expected filename (may vary)

// EnsureModelExists checks if the word embedding model is available locally.
// If not found, it initiates the download process with a progress UI.
// The model is downloaded once and cached for future use.
//
// Returns an error if download or extraction fails.
func EnsureModelExists() error {
	// Check if model already exists locally
	if _, err := os.Stat("model.vec"); err == nil {
		return nil // Model already available
	}

	// Start download process with progress UI
	p := tea.NewProgram(newDownloadModel())
	if _, err := p.Run(); err != nil {
		return err
	}

	// Extract and prepare the model file
	fmt.Println("\nExtracting model...")
	return unzipAndPrepare()
}

// unzipAndPrepare extracts the word vectors file from the downloaded zip
// and renames it to the standard "model.vec" filename for consistency.
func unzipAndPrepare() error {
	// Extract all files from the zip archive
	if err := unzip(zipFile, "."); err != nil {
		return fmt.Errorf("failed to unzip: %v", err)
	}

	// Rename the extracted file to standard name
	fmt.Printf("Renaming %s to model.vec\n", targetFile)
	err := os.Rename(targetFile, "model.vec")
	if err != nil {
		return fmt.Errorf("failed to rename %s: %v", targetFile, err)
	}

	// Clean up: remove the original zip file
	os.Remove(zipFile)
	return nil
}

// unzip extracts a zip archive to the specified destination directory.
// It automatically detects .txt files (word vector files) and extracts them.
// Updates the global targetFile variable with the actual filename found.
func unzip(src, dest string) error {
	// Open the zip archive for reading
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// List all files in the archive for debugging
	fmt.Println("Files in zip:")
	for _, f := range r.File {
		fmt.Printf("  %s\n", f.Name)
	}

	// Look for any .txt file (word vectors are typically in text format)
	for _, f := range r.File {
		if filepath.Ext(f.Name) == ".txt" {
			fmt.Printf("Extracting: %s\n", f.Name)

			// Open the file within the archive
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			// Create the output file in destination directory
			outFile, err := os.Create(filepath.Join(dest, f.Name))
			if err != nil {
				return err
			}
			defer outFile.Close()

			// Copy file contents from archive to disk
			_, err = io.Copy(outFile, rc)
			if err != nil {
				return err
			}

			// Update the global variable to match the actual filename
			targetFile = f.Name
			return nil // Successfully extracted the word vectors file
		}
	}

	return fmt.Errorf("no .txt file found in zip archive")
}
