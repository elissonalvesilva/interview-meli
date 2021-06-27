package services

import (
	"github.com/elissonalvesilva/interview-meli/analyzer/shared"
	"github.com/elissonalvesilva/interview-meli/analyzer/tests/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMatrixFromDNAList(t *testing.T) {
	t.Run("Should return a error if dna is not provided", func(t *testing.T) {
		t.Parallel()
		resp, err := GetMatrixFromDNAList(nil)
		assert.Nil(t, resp)
		assert.Error(t, err, ErrDNAEmpty)
	})

	t.Run("Should return a error if dna slice is not square", func(t *testing.T) {
		t.Parallel()
		resp, err := GetMatrixFromDNAList(mock.DNAInvalidMatrix)
		assert.Nil(t, resp)
		assert.Error(t, err, ErrInvalidMatrix)
	})

	t.Run("Should return a error if dna slice contains a invalid character (permitted chars: G, T, A, C)", func(t *testing.T) {
		t.Parallel()
		resp, err := GetMatrixFromDNAList(mock.DNAInvalidChar)
		assert.Nil(t, resp)
		assert.Error(t, err, ErrInvalidCharMatrix)
	})

	t.Run("Should return a valid matrix", func(t *testing.T) {
		t.Parallel()
		resp, err := GetMatrixFromDNAList(mock.DNAValid)
		assert.NotNil(t, resp)
		assert.Nil(t, err)
		assert.Equal(t, mock.DNAMatrixValid, resp)
	})
}

func TestIsValidMatrix(t *testing.T) {
	t.Run("Should return false if is a invalid matrix", func(t *testing.T) {
		t.Parallel()

		resp := IsValidMatrix(mock.DNAInvalidMatrix)
		assert.False(t, resp)
	})

	t.Run("Should return true if is a valid matrix", func(t *testing.T) {
		t.Parallel()

		resp := IsValidMatrix(mock.DNAValid)
		assert.True(t, resp)
	})
}

func TestSearchForSimianValidRepetitions(t *testing.T) {
	t.Run("Should return true if countLettersFound is equals to 4", func(t *testing.T) {
		t.Parallel()

		resp := SearchForSimianValidRepetitions(mock.DNAMatrixValid, 1,1,1,1,4, "T")
		assert.NotNil(t, resp)
		assert.True(t, resp)
	})

	t.Run("Should return false if has a invalid coordinate or current letter is not equals to next letter", func(t *testing.T) {
		t.Parallel()

		MatrixLength = len(mock.DNAMatrixValid)

		resp := SearchForSimianValidRepetitions(mock.DNAMatrixValid, 1,1,1,1,1, "C")
		assert.NotNil(t, resp)
		assert.False(t, resp)
	})

	t.Run("Should return false if countLettersFound is less than 0", func(t *testing.T) {
		t.Parallel()

		MatrixLength = len(mock.DNAMatrixValid)

		resp := SearchForSimianValidRepetitions(mock.DNAMatrixValid, 1,1,1,1,0, "T")
		assert.NotNil(t, resp)
		assert.False(t, resp)
	})
}

func TestDNAAnalyzer(t *testing.T) {
	t.Run("Should return a error if dna is empty", func(t *testing.T) {
		t.Parallel()

		sut := NewAnalyzer()

		resp, err := sut.DNAAnalyzer(nil)
		assert.Equal(t, shared.EmptyDNA, resp)
		assert.Error(t, err, ErrDNAEmpty)
	})

	t.Run("Should return DNA human type", func(t *testing.T) {
		t.Parallel()

		sut := NewAnalyzer()

		resp, err := sut.DNAAnalyzer(mock.DNAHuman)
		assert.Nil(t, err)
		assert.Equal(t, shared.DNATypeHuman, resp)
	})

	t.Run("Should return DNA simian type", func(t *testing.T) {
		t.Parallel()

		sut := NewAnalyzer()

		resp, err := sut.DNAAnalyzer(mock.DNAValid)
		assert.Nil(t, err)
		assert.Equal(t, shared.DNATypeSimian, resp)
	})
}