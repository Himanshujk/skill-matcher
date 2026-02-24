// Package main provides data indexing functionality for caching search results
// and pre-processed Excel data. Uses Go's gob encoding for efficient binary serialization.
package main

import (
	"encoding/gob"
	"os"
)

// IndexedRow represents a single row from an Excel file along with its metadata.
// This structure is used for caching processed data to avoid re-parsing large files.
//
// Example usage:
//
//	row := IndexedRow{
//	  File: "resume_database.xlsx",
//	  Row:  ["John Doe", "Python, JavaScript", "Computer Science"],
//	  Vec:  [0.1, -0.2, 0.3, ...],  // Pre-computed embedding vector
//	}
type IndexedRow struct {
	File string    // Source Excel file path
	Row  []string  // Excel row data (all columns)
	Vec  []float64 // Pre-computed semantic vector for the row
}

// SaveIndex serializes indexed rows to a binary file using gob encoding.
// This allows for fast loading of pre-processed Excel data in future searches.
//
// Performance note: Binary storage is much faster than parsing Excel files repeatedly.
func SaveIndex(path string, rows []IndexedRow) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	return enc.Encode(rows)
}

// LoadIndex deserializes previously saved indexed rows from a binary file.
// Returns an error if the file doesn't exist or is corrupted.
//
// Usage:
//
//	cachedRows, err := LoadIndex("search_cache.gob")
//	if err == nil {
//	  // Use cached data instead of parsing Excel files
//	}
func LoadIndex(path string) ([]IndexedRow, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows []IndexedRow
	dec := gob.NewDecoder(file)
	err = dec.Decode(&rows)
	return rows, err
}
