package services

import (
	"errors"
	"github.com/elissonalvesilva/interview-meli/api/domain/entity"
	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
	"github.com/elissonalvesilva/interview-meli/api/shared/constants"
	"github.com/elissonalvesilva/interview-meli/api/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStatsSimianHuman(t *testing.T) {
	t.Run("Should return a error if StatsHumanSimian throws", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDNARepository := mock.NewMockDNARepository(ctrl)

		handlerError := "error to get stats"

		mockDNARepository.EXPECT().StatsHumanSimian().Return(protocols.StatsResponse{}, errors.New(handlerError))

		sut := NewDNAService(mockDNARepository)
		resp, err := sut.GetStatsSimianHuman()
		assert.NotNil(t, err, handlerError)
		assert.Error(t, err, handlerError)
		assert.Equal(t, protocols.StatsResponse{}, resp)
	})

	t.Run("Should return stats from StatsHumanSimian", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDNARepository := mock.NewMockDNARepository(ctrl)


		mockDNARepository.EXPECT().StatsHumanSimian().Return(mock.Stats, nil)

		sut := NewDNAService(mockDNARepository)
		resp, err := sut.GetStatsSimianHuman()
		assert.Nil(t, err)
		assert.Equal(t, mock.Stats, resp)
	})
}

func TestAddDNA(t *testing.T) {
	t.Run("Should return a error if entity.NewDNA throws", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDNARepository := mock.NewMockDNARepository(ctrl)

		sut := NewDNAService(mockDNARepository)
		_, err := sut.CreateDNA([]string{})

		assert.NotNil(t, err, entity.ErrEmptyDNA)
		assert.Error(t, err, entity.ErrEmptyDNA)
	})

	t.Run("Should return a error if CheckIfDNAExists throws", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDNARepository := mock.NewMockDNARepository(ctrl)

		handlerError := "handler error"
		mockDNARepository.EXPECT().CheckIfDNAExists(mock.DNA).Return(false, errors.New("handlerError"))

		sut := NewDNAService(mockDNARepository)
		resp, err := sut.CreateDNA(mock.DNA)

		assert.NotNil(t, err, handlerError)
		assert.Error(t, err, handlerError)
		assert.Equal(t, entity.DNASequence{}, resp)
	})

	t.Run("Should return a error if dna exists in database", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDNARepository := mock.NewMockDNARepository(ctrl)

		errorDNA := errors.New(constants.DNAExists)
		mockDNARepository.EXPECT().CheckIfDNAExists(mock.DNA).Return(true, nil)

		sut := NewDNAService(mockDNARepository)
		resp, err := sut.CreateDNA(mock.DNA)

		assert.NotNil(t, err, errorDNA)
		assert.Error(t, err, errorDNA)
		assert.Equal(t, entity.DNASequence{}, resp)
	})

	t.Run("Should return a error if AddDNA throws", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDNARepository := mock.NewMockDNARepository(ctrl)

		handlerError := "handlerError"
		mockDNARepository.EXPECT().CheckIfDNAExists(mock.DNA).Return(false, nil)
		mockDNARepository.EXPECT().AddDNA(mock.DNA, "SIMIAN").Return(errors.New(handlerError))

		sut := NewDNAService(mockDNARepository)
		resp, err := sut.CreateDNA(mock.DNA)

		assert.NotNil(t, err, handlerError)
		assert.Error(t, err, handlerError)
		assert.Equal(t, entity.DNASequence{}, resp)
	})

	t.Run("Should return a dna on success", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDNARepository := mock.NewMockDNARepository(ctrl)

		mockDNARepository.EXPECT().CheckIfDNAExists(mock.DNA).Return(false, nil)
		mockDNARepository.EXPECT().AddDNA(mock.DNA, "SIMIAN").Return(nil)

		sut := NewDNAService(mockDNARepository)
		resp, err := sut.CreateDNA(mock.DNA)

		assert.Nil(t, err)
		assert.Equal(t, mock.AddedDNA, resp)
	})
}