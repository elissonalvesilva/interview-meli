package analyzer_grpc_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/elissonalvesilva/interview-meli/pkg/proto/dna_pb"
	"google.golang.org/grpc"
	"log"
	"os"
)

type AnalyzerGRPCService struct {}

func NewAnalyzerService() *AnalyzerGRPCService {
	return &AnalyzerGRPCService{}
}

func (grpcService *AnalyzerGRPCService) AnalyzeDNA(dna []string) (string, error) {
	address := fmt.Sprintf("%s:%s", os.Getenv("GRPC_ANALYZER_HOST"), os.Getenv("GRPC_ANALYZER_PORT"))
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer connection.Close()

	client := dna_pb.NewDNAServiceAnalyzerClient(connection)

	request := &dna_pb.DNARequest{
		Dna: dna,
	}

	res, err := client.Analyze(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during execution: %v", err)
	}

	typeDNA := res.TypeDNA
	errAnalyzer := res.DnaError

	if res.DnaError == "" {
		return typeDNA, nil
	}

	return typeDNA, errors.New(errAnalyzer)
}
