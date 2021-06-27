package services

import (
	"github.com/elissonalvesilva/interview-meli/api/domain/services"
	"github.com/elissonalvesilva/interview-meli/api/infra/db/mongodb"
	"github.com/elissonalvesilva/interview-meli/api/infra/db/mongodb/helper"
)

func MakeDNAService(db helper.MongoDB) *services.DNAService {
	repository := mongodb.NewDNADatabase(db)
	service := services.NewDNAService(repository)

	return service
}
