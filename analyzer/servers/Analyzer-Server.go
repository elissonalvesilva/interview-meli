package servers

import (
	"context"
	"github.com/elissonalvesilva/interview-meli/analyzer/domain/services"
	"github.com/elissonalvesilva/interview-meli/analyzer/presenters/controllers"
	"github.com/elissonalvesilva/interview-meli/pkg/proto/dna_pb"
)

type AnalyzerServer struct {}

func NewAnalyzerServer() *AnalyzerServer {
	return &AnalyzerServer{}
}

func (a *AnalyzerServer) Analyze(ctx context.Context, request *dna_pb.DNARequest) (*dna_pb.DNAResponse, error) {
	service := services.NewAnalyzer()
	controller := controllers.NewAnalyzerController(*service)

	typeDna, err := controller.AnalyzeDNA(request.GetDna())

	if err != nil {
		return &dna_pb.DNAResponse{
			TypeDNA: typeDna,
			DnaError: err.Error(),
		}, err
	}

	return &dna_pb.DNAResponse{
		TypeDNA: typeDna,
		DnaError: "",
	}, nil
}
