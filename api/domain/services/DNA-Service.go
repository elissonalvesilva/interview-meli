package services

import (
	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
	"github.com/elissonalvesilva/interview-meli/api/domain/repository"
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
