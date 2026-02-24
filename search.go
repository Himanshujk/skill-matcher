// Package main provides mathematical functions for vector operations
// and text processing utilities used in semantic similarity calculations.
package main

import (
	"math"
	"strings"
)

// cosine calculates the cosine similarity between two vectors.
// Cosine similarity measures the cosine of the angle between two vectors,
// returning values from -1 (opposite directions) to 1 (same direction).
// Values closer to 1 indicate higher similarity.
//
// Formula: cos(θ) = (a·b) / (||a|| * ||b||)
// where a·b is dot product and ||a|| is magnitude of vector a
func cosine(a, b []float64) float64 {
	// Calculate dot product and vector magnitudes in a single pass
	var dot, normA, normB float64
	for i := range a {
		dot += a[i] * b[i]   // Dot product: sum of element-wise multiplication
		normA += a[i] * a[i] // Squared magnitude of vector a
		normB += b[i] * b[i] // Squared magnitude of vector b
	}
	// Return cosine similarity: dot product / (magnitude_a * magnitude_b)
	return dot / (math.Sqrt(normA) * math.Sqrt(normB))
}

// averageVec computes the element-wise average of multiple vectors.
// This is used to create a single representative vector from multiple skill embeddings.
// For example, if searching for ["Python", "JavaScript"], this function
// creates one vector that represents both skills combined.
//
// Returns nil if no vectors are provided.
func averageVec(vecs [][]float64) []float64 {
	if len(vecs) == 0 {
		return nil
	}
	// Initialize output vector with same dimensions as input vectors
	dim := len(vecs[0])
	out := make([]float64, dim)

	// Sum all vectors element-wise
	for _, v := range vecs {
		for i := range v {
			out[i] += v[i]
		}
	}

	// Divide by vector count to get average
	for i := range out {
		out[i] /= float64(len(vecs))
	}
	return out
}

// parseSkills takes a comma-separated string of skills and returns a cleaned slice.
// It handles common input variations:
// - Converts to lowercase for consistent matching
// - Trims whitespace from each skill
// - Filters out empty strings
// - Splits on commas
//
// Example: "Python, JavaScript ,  data science" → ["python", "javascript", "data science"]
func parseSkills(input string) []string {
	parts := strings.Split(input, ",")
	var out []string
	for _, p := range parts {
		// Clean each skill: trim whitespace and convert to lowercase
		s := strings.TrimSpace(strings.ToLower(p))
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
