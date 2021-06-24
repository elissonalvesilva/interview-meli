package mock

import (
	"github.com/elissonalvesilva/interview-meli/api/domain/entity"
	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
)

var Stats = protocols.StatsResponse{
	CountHumanDNA: 100,
	CountSimianDNA: 101,
	Ratio: 1.01,
}

var DNA = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

var AddedDNA = entity.DNASequence{
	DNA: DNA,
}

