package controllers

import (
	"github.com/elissonalvesilva/interview-meli/api/infra/db/mongodb/helper"
	"github.com/elissonalvesilva/interview-meli/api/presenters/controllers"
	"github.com/elissonalvesilva/interview-meli/api/server/factories/services"
)

func MakeDNAController(db helper.MongoDB) *controllers.DNAController {
	service := services.MakeDNAService(db)
	controller := controllers.NewDNAController(*service)
	return controller
}
