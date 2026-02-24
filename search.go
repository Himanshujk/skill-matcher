package main

import (
	"math"
	"strings"
)

func cosine(a, b []float64) float64 {
	var dot, normA, normB float64
	for i := range a {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	return dot / (math.Sqrt(normA) * math.Sqrt(normB))
}

func averageVec(vecs [][]float64) []float64 {
	if len(vecs) == 0 {
		return nil
	}
	dim := len(vecs[0])
	out := make([]float64, dim)
	for _, v := range vecs {
		for i := range v {
			out[i] += v[i]
		}
	}
	for i := range out {
		out[i] /= float64(len(vecs))
	}
	return out
}

func parseSkills(input string) []string {
	parts := strings.Split(input, ",")
	var out []string
	for _, p := range parts {
		s := strings.TrimSpace(strings.ToLower(p))
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
