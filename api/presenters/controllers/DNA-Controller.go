package controllers

import (
	"encoding/json"
	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
	"github.com/elissonalvesilva/interview-meli/api/domain/services"
	"net/http"
)

type DNAController struct {
	service services.DNAService
}

func NewDNAController(service services.DNAService) *DNAController {
	return &DNAController{
		service: service,
	}
}

func (ctrl *DNAController) AnalyzeDNA(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var rd protocols.ReceiveDNA
	errDecode := decoder.Decode(&rd)
	if errDecode != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Must be pass a dna param")
		return
	}

	_, errService := ctrl.service.CreateDNA(rd.DNA)

	if errService != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ctrl *DNAController) GetStatsDNAs(w http.ResponseWriter, r *http.Request) {
	stats, errService := ctrl.service.GetStatsSimianHuman()

	if errService != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stats)
}
