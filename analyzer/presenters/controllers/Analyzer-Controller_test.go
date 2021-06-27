package controllers

import (
	"github.com/elissonalvesilva/interview-meli/analyzer/domain/services"
	"github.com/elissonalvesilva/interview-meli/analyzer/tests/mock"
	"github.com/elissonalvesilva/interview-meli/api/shared/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnalyzerController(t *testing.T) {
	t.Run("Should return empty type DNA and a error", func(t *testing.T) {
		t.Parallel()

		service := services.NewAnalyzer()
		sut := NewAnalyzerController(*service)

		typeDNA, err := sut.AnalyzeDNA(nil)
		assert.Error(t, err, services.ErrDNAEmpty)
		assert.Equal(t, "", typeDNA)
	})

	t.Run("Should return a type dna and error equals to nil", func(t *testing.T) {
		t.Parallel()

		service := services.NewAnalyzer()
		sut := NewAnalyzerController(*service)

		typeDNA, err := sut.AnalyzeDNA(mock.DNAValid)
		assert.Nil(t, err)
		assert.Equal(t, constants.DNATypeSimian, typeDNA)
	})
}
