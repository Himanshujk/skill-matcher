package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

const modelURL = "https://nlp.stanford.edu/data/wordvecs/glove.2024.wikigiga.100d.zip"
const zipFile = "glove.2024.wikigiga.100d.zip"

var targetFile = "glove.2024.wikigiga.100d.txt" // will be updated when we find the actual file

func EnsureModelExists() error {
	if _, err := os.Stat("model.vec"); err == nil {
		return nil
	}

	p := tea.NewProgram(newDownloadModel())
	if _, err := p.Run(); err != nil {
		return err
	}

	fmt.Println("\nExtracting model...")
	return unzipAndPrepare()
}

func unzipAndPrepare() error {
	if err := unzip(zipFile, "."); err != nil {
		return fmt.Errorf("failed to unzip: %v", err)
	}

	fmt.Printf("Renaming %s to model.vec\n", targetFile)
	err := os.Rename(targetFile, "model.vec")
	if err != nil {
		return fmt.Errorf("failed to rename %s: %v", targetFile, err)
	}

	os.Remove(zipFile)
	return nil
}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	// First, let's see what files are in the zip
	fmt.Println("Files in zip:")
	for _, f := range r.File {
		fmt.Printf("  %s\n", f.Name)
	}

	// Look for any .txt file (word vectors file)
	for _, f := range r.File {
		if filepath.Ext(f.Name) == ".txt" {
			fmt.Printf("Extracting: %s\n", f.Name)

			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			// Extract with the original name first, we'll rename it later
			outFile, err := os.Create(filepath.Join(dest, f.Name))
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return err
			}

			// Update the global targetFile to match what we actually found
			targetFile = f.Name
			return nil
		}
	}

	return fmt.Errorf("no .txt file found in zip archive")
}
