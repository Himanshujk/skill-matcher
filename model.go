// Package main provides functionality for loading and using pre-trained word embeddings.
// It supports GloVe (Global Vectors for Word Representation) format models
// for semantic similarity calculations in skill matching.
package main

import (
	"bufio"
	"os"
	"skillsearch/helpers"
	"strconv"
	"strings"
)

// Model represents a loaded word embedding model.
// It contains a map of words to their vector representations and the dimensionality.
//
// Example:
//
//	model := LoadModel("glove.txt")
//	vector := model.Embed("python programming")
type Model struct {
	Vectors map[string][]float64 // Word-to-vector mappings
	Dim     int                  // Vector dimensionality (e.g., 100, 200, 300)
}

// LoadModel reads a GloVe format text file and constructs a word embedding model.
// The file format expected is:
//
//	word float1 float2 float3 ... floatN
//
// For example:
//
//	python 0.1 -0.2 0.3 0.4 ...
//	programming -0.1 0.5 -0.3 0.2 ...
//
// The function ensures dimensional consistency by setting the dimension
// from the first valid vector and skipping vectors with different dimensions.
func LoadModel(path string) (*Model, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize model with empty vector map
	model := &Model{Vectors: make(map[string][]float64)}
	scanner := bufio.NewScanner(file)

	// Track first line to establish vector dimensions
	firstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		// Skip lines with insufficient data (word + at least 10 vector components)
		if len(parts) < 10 {
			continue
		}

		// Extract word (first field) and vector components (remaining fields)
		word := parts[0]
		vecLen := len(parts) - 1
		vec := make([]float64, vecLen)

		// Parse floating point values for vector components
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

		// Store word-vector mapping
		model.Vectors[word] = vec
	}

	return model, nil
}

// Embed converts text into a vector representation by averaging word embeddings.
// The process:
// 1. Splits text into individual words
// 2. Looks up each word in the embedding model
// 3. Averages all found word vectors
// 4. Returns nil if no words have embeddings
//
// Example:
//
//	vector := model.Embed("python programming")  // averages "python" + "programming" vectors
//	vector := model.Embed("unknown_word")        // returns nil
func (m *Model) Embed(text string) []float64 {
	// Tokenize text into lowercase words
	words := make([]string, 0)
	for _, s := range strings.Split(text, ",") {
		n := helpers.NormalizeSkill(s)
		if n != "" {
			words = append(words, n)
		}
	}
	// Initialize output vector with model dimensions
	vec := make([]float64, m.Dim)

	// Track how many words had embeddings for averaging
	count := 0
	for _, w := range words {
		if v, ok := m.Vectors[w]; ok {
			// Add bounds checking to prevent index out of range
			maxLen := len(v)
			if maxLen > m.Dim {
				maxLen = m.Dim
			}

			// Add word vector to running sum
			for i := 0; i < maxLen; i++ {
				vec[i] += v[i]
			}
			count++
		}
	}

	// Return nil if no words had embeddings
	if count == 0 {
		return nil
	}

	// Average the sum by dividing by word count
	for i := range vec {
		vec[i] /= float64(count)
	}
	return vec
}
