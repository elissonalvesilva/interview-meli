package services

import (
	"errors"
	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
	"github.com/elissonalvesilva/interview-meli/api/domain/services"
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

		sut := services.NewDNAService(mockDNARepository)
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

		sut := services.NewDNAService(mockDNARepository)
		resp, err := sut.GetStatsSimianHuman()
		assert.Nil(t, err)
		assert.Equal(t, mock.Stats, resp)
	})
}
