package repository

import (
	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
)

type DNARepository interface {
	CheckIfDNAExists([]string) (bool, error)
	AddDNA([]string, string) error
	StatsHumanSimian() (protocols.StatsResponse, error)
}
