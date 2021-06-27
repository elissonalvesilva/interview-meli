package services

import (
	"github.com/elissonalvesilva/interview-meli/api/domain/services"
	"github.com/elissonalvesilva/interview-meli/api/infra/db/mongodb"
	"github.com/elissonalvesilva/interview-meli/api/infra/db/mongodb/helper"
	analyzer_grpc_client "github.com/elissonalvesilva/interview-meli/api/infra/grpc/analyzer-grpc-client"
)

func MakeDNAService(db helper.MongoDB) *services.DNAService {
	repository := mongodb.NewDNADatabase(db)
	serviceGRPC := analyzer_grpc_client.NewAnalyzerService()
	analyzerService := services.NewDNAAnalyzeService(serviceGRPC)
	service := services.NewDNAService(repository, *analyzerService)

	return service
}
