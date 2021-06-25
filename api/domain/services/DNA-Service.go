package services

import (
	"errors"
	"github.com/elissonalvesilva/interview-meli/api/domain/entity"
	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
	"github.com/elissonalvesilva/interview-meli/api/domain/repository"
	"github.com/elissonalvesilva/interview-meli/api/shared/constants"
)

type DNAService struct {
	repo repository.DNARepository
}

func NewDNAService(repo repository.DNARepository) *DNAService {
	return &DNAService{
		repo: repo,
	}
}

func (s *DNAService) GetStatsSimianHuman() (protocols.StatsResponse, error) {
	stats, err := s.repo.StatsHumanSimian()

	if err != nil {
		return protocols.StatsResponse{}, err
	}

	return stats, nil
}

func (s *DNAService) CreateDNA(dna []string) (entity.DNASequence, error) {
	createdDNA, errCreateDNA := entity.NewDNA(dna)
	if errCreateDNA != nil {
		return entity.DNASequence{}, errCreateDNA
	}

	alreadyInDB, err := s.repo.CheckIfDNAExists(createdDNA.DNA)
	if err != nil {
		return entity.DNASequence{}, err
	}

	if alreadyInDB {
		return entity.DNASequence{}, errors.New(constants.DNAExists)
	}

	dnaType := "SIMIAN"
	errToAddDNA := s.repo.AddDNA(createdDNA.DNA, dnaType)
	if errToAddDNA != nil {
		return entity.DNASequence{}, errToAddDNA
	}

	return *createdDNA, nil
}
