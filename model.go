package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Model struct {
	Vectors map[string][]float64
	Dim     int
}

func LoadModel(path string) (*Model, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	model := &Model{Vectors: make(map[string][]float64)}
	scanner := bufio.NewScanner(file)

	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 10 {
			continue
		}

		word := parts[0]
		vecLen := len(parts) - 1
		vec := make([]float64, vecLen)

		for i := 1; i < len(parts); i++ {
			val, _ := strconv.ParseFloat(parts[i], 64)
			vec[i-1] = val
		}

		// Set dimension from the first valid vector and ensure consistency
		if firstLine {
			model.Dim = vecLen
			firstLine = false
		} else if vecLen != model.Dim {
			// Skip vectors with inconsistent dimensions
			continue
		}

		model.Vectors[word] = vec
	}

	return model, nil
}

func (m *Model) Embed(text string) []float64 {
	words := strings.Fields(strings.ToLower(text))
	vec := make([]float64, m.Dim)

	count := 0
	for _, w := range words {
		if v, ok := m.Vectors[w]; ok {
			// Add bounds checking to prevent index out of range
			maxLen := len(v)
			if maxLen > m.Dim {
				maxLen = m.Dim
			}

			for i := 0; i < maxLen; i++ {
				vec[i] += v[i]
			}
			count++
		}
	}

	if count == 0 {
		return nil
	}

	for i := range vec {
		vec[i] /= float64(count)
	}
	return vec
}
