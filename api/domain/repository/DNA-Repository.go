package repository

import "github.com/elissonalvesilva/interview-meli/api/domain/protocols"

type DNARepository interface {
	StatsHumanSimian() (protocols.StatsResponse, error)
}
