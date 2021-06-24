package main

import (
	"testing"
)

func benchmarkIsSimian(dna []string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		isSimian(dna)
	}
}

func BenchmarkIsSimian1(b *testing.B)  {
	dna := []string{
		"CTGAGA",
		"CTGAGC",
		"TATTGT",
		"AGAGGG",
		"CCCCTG",
		"TCACTG",
	}
	benchmarkIsSimian(dna, b)
}

func BenchmarkIsSimian2(b *testing.B)  {
	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG",
	}
	benchmarkIsSimian(dna, b)
}
