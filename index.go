package main

import (
	"encoding/gob"
	"os"
)

type IndexedRow struct {
	File string
	Row  []string
	Vec  []float64
}

func SaveIndex(path string, rows []IndexedRow) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	return enc.Encode(rows)
}

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
