package mongodb

import (
	"context"
	"errors"
	"github.com/elissonalvesilva/interview-meli/api/domain/entity"
	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
	"github.com/elissonalvesilva/interview-meli/api/tests/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"

	"testing"
)

func TestCheckIfDNAExists(t *testing.T) {
	t.Run("Should return a error if FindOne Throws", func(t *testing.T) {
		t.Parallel()

		handlerError := "error"
		ctrl := gomock.NewController(t)
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
		mockSingleResult := mock.NewMockSingleResultHelper(ctrl)

		mockDb.EXPECT().Collection("dna").Return(mockCollection)
		mockSingleResult.EXPECT().Decode(gomock.Any()).Return(errors.New(handlerError))

		mockCollection.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(mockSingleResult)

		sut := NewDNADatabase(mockDb)
		resp, err := sut.CheckIfDNAExists(mock.DNA)
		assert.Error(t, err, handlerError)
		assert.Equal(t, false, resp)
	})

	t.Run("Should return a error if FindOne Throws with ErrNoDocuments", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
		mockSingleResult := mock.NewMockSingleResultHelper(ctrl)

		mockDb.EXPECT().Collection("dna").Return(mockCollection)
		mockSingleResult.EXPECT().Decode(gomock.Any()).Return(errors.New(mongo.ErrNoDocuments.Error()))

		mockCollection.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(mockSingleResult)

		sut := NewDNADatabase(mockDb)
		resp, err := sut.CheckIfDNAExists(mock.DNA)
		assert.Error(t, err, mongo.ErrNoDocuments)
		assert.Equal(t, false, resp)
	})

	t.Run("Should return true if exists dna in database", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
		mockSingleResult := mock.NewMockSingleResultHelper(ctrl)

		mockDb.EXPECT().Collection("dna").Return(mockCollection)
		mockSingleResult.EXPECT().Decode(gomock.Any()).Return(nil)

		mockCollection.EXPECT().FindOne(gomock.Any(), gomock.Any()).Return(mockSingleResult)

		sut := NewDNADatabase(mockDb)
		resp, err := sut.CheckIfDNAExists(mock.DNA)
		assert.Nil(t, err)
		assert.Equal(t, true, resp)
	})
}

func TestAddDNA(t *testing.T) {
	t.Run("Should return a error if InsertOne throws", func(t *testing.T) {
		t.Parallel()

		handlerError := "error"
		ctrl := gomock.NewController(t)
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
		mockDb.EXPECT().Collection("dna").Return(mockCollection)

		mockCollection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return(entity.DNASequence{}, errors.New(handlerError))

		sut := NewDNADatabase(mockDb)

		err := sut.AddDNA(mock.DNA, "SIMIAN")
		assert.Error(t, err, handlerError)
	})

	t.Run("Should return nil if insert dna in database", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
		mockDb.EXPECT().Collection("dna").Return(mockCollection)

		mockCollection.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return(mock.AddedDNA, nil)

		sut := NewDNADatabase(mockDb)

		err := sut.AddDNA(mock.DNA, "SIMIAN")
		assert.Nil(t, err)
	})
}

func TestStatsHumanSimian(t *testing.T)  {
	t.Run("Should return a error if CountDocuments on query human throws", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		handlerError := "error count human"
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
		mockDb.EXPECT().Collection("dna").Return(mockCollection)

		ctx := context.Background()

		mockCollection.EXPECT().CountDocuments(ctx, gomock.Any()).Return(int64(0), errors.New(handlerError))

		sut := NewDNADatabase(mockDb)

		resp, err := sut.StatsHumanSimian()
		assert.Error(t, err, handlerError)
		assert.Equal(t, protocols.StatsResponse{}, resp)
	})

	t.Run("Should return a error if CountDocuments on query simians throws", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		handlerError := "error count simian"
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
		mockDb.EXPECT().Collection("dna").Return(mockCollection)

		ctx := context.Background()

		countHuman := mock.Stats.CountHumanDNA
		gomock.InOrder(
			mockCollection.EXPECT().CountDocuments(ctx, gomock.Any()).Return(countHuman, nil),
			mockCollection.EXPECT().CountDocuments(ctx, gomock.Any()).Return(int64(0), errors.New(handlerError)),
		)

		sut := NewDNADatabase(mockDb)

		resp, err := sut.StatsHumanSimian()
		assert.Error(t, err, handlerError)
		assert.Equal(t, protocols.StatsResponse{}, resp)
	})

	t.Run("Should return a ratio equals to 0 if count human dna is equals to 0", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		//handlerError := "error"
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
		mockDb.EXPECT().Collection("dna").Return(mockCollection)

		ctx := context.Background()

		countHuman := int64(0)
		countSimian := mock.Stats.CountSimianDNA
		gomock.InOrder(
			mockCollection.EXPECT().CountDocuments(ctx, gomock.Any()).Return(countHuman, nil),
			mockCollection.EXPECT().CountDocuments(ctx, gomock.Any()).Return(countSimian, nil),
		)

		sut := NewDNADatabase(mockDb)

		resp, err := sut.StatsHumanSimian()
		assert.Nil(t, err)
		assert.Equal(t, protocols.StatsResponse{
			CountHumanDNA: 0,
			CountSimianDNA: countSimian,
			Ratio: 0,
		}, resp)
	})

	t.Run("Should return a response with ratio", func(t *testing.T) {
		t.Parallel()

		ctrl := gomock.NewController(t)

		//handlerError := "error"
		mockDb := mock.NewMockMongoDB(ctrl)
		mockCollection := mock.NewMockCollectionHelper(ctrl)
		mockDb.EXPECT().Collection("dna").Return(mockCollection)

		ctx := context.Background()

		countHuman := mock.Stats.CountHumanDNA
		countSimian := mock.Stats.CountSimianDNA
		gomock.InOrder(
			mockCollection.EXPECT().CountDocuments(ctx, gomock.Any()).Return(countHuman, nil),
			mockCollection.EXPECT().CountDocuments(ctx, gomock.Any()).Return(countSimian, nil),
		)

		sut := NewDNADatabase(mockDb)

		resp, err := sut.StatsHumanSimian()
		assert.Nil(t, err)
		assert.Equal(t, mock.Stats, resp)
	})
}
