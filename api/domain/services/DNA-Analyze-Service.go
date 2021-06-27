package services

import "github.com/elissonalvesilva/interview-meli/api/domain/repository"

type DNAAnalyzeService struct {
	grpcClientAnalyze repository.AnalyzeGRPC
}

func NewDNAAnalyzeService(grpcClientAnalyze repository.AnalyzeGRPC) *DNAAnalyzeService {
	return &DNAAnalyzeService{
		grpcClientAnalyze: grpcClientAnalyze,
	}
}

func (client *DNAAnalyzeService) Analyze(dna []string) (string, error) {
	analyzedDNA, err := client.grpcClientAnalyze.AnalyzeDNA(dna)
	if err != nil {
		return analyzedDNA, err
	}

	return analyzedDNA, nil
}
