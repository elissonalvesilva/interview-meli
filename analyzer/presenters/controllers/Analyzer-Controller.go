package controllers

import "github.com/elissonalvesilva/interview-meli/analyzer/domain/services"

type AnalyzerController struct {
	service services.Analyzer
}

func NewAnalyzerController(service services.Analyzer) *AnalyzerController {
	return &AnalyzerController{
		service: service,
	}
}

func (s *AnalyzerController) AnalyzeDNA(dna []string) (string, error) {
	typeDNA, err := s.service.DNAAnalyzer(dna)
	return typeDNA, err
}
