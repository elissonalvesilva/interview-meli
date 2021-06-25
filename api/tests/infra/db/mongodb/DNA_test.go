package mongodb

import (
	"github.com/elissonalvesilva/interview-meli/api/tests/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestStatsHumanSimian(t *testing.T)  {
	t.Run("Test", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		t.Parallel()

		handlerError := "error"
		ctrl := gomock.NewController(t)
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
	})
}
