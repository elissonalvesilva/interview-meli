package mongodb

import (
	"context"
	"github.com/elissonalvesilva/interview-meli/api/domain/entity"
	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
	"github.com/elissonalvesilva/interview-meli/api/infra/db/mongodb/helper"
	"github.com/elissonalvesilva/interview-meli/api/shared/constants"
	"go.mongodb.org/mongo-driver/bson"
)

type DNADatabase struct {
	db helper.MongoDB
}

const collectionName = "dna"

func NewDNADatabase (db helper.MongoDB) *DNADatabase {
	return &DNADatabase{
		db: db,
	}
}

type dnaBson struct {
	DNA []string `json:"dna"`
	Type string `json:"type"`
}

func FromDNA(dna []string, dnaType string) *dnaBson {
	return &dnaBson{
		DNA: dna,
		Type: dnaType,
	}
}

func (d *DNADatabase) CheckIfDNAExists(dna []string) (bool, error) {
	dnaCollection := d.db.Collection(collectionName)
	var dnaEntity entity.DNASequence

	var (
		query = bson.M{"dna": dna}
	)
	ctx := context.Background()

	err := dnaCollection.FindOne(ctx, query).Decode(&dnaEntity)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (d *DNADatabase) AddDNA(dna []string, dnaType string) error {
	dnaCollection := d.db.Collection(collectionName)

	bson := FromDNA(dna, dnaType)
	ctx := context.Background()
	_, err := dnaCollection.InsertOne(ctx, bson)

	if err != nil {
		return err
	}
	return nil

}

func (d *DNADatabase) StatsHumanSimian() (protocols.StatsResponse, error) {
	dnaCollection := d.db.Collection(collectionName)
	ctx := context.Background()

	var (
		queryHuman = bson.M{"type" : constants.DNATypeHuman}
		querySimian = bson.M{"type" : constants.DNATypeSimian}
	)

	countHuman, errHuman := dnaCollection.CountDocuments(ctx, queryHuman)
	if errHuman != nil {
		return protocols.StatsResponse{}, errHuman
	}

	countSimian, errSimian := dnaCollection.CountDocuments(ctx, querySimian)

	if errSimian != nil {
		return protocols.StatsResponse{}, errSimian
	}

	if countHuman == 0 {
		return protocols.StatsResponse {
			CountHumanDNA: 0,
			CountSimianDNA: countSimian,
			Ratio: 0,
		}, nil
	}

	ratio := float64(countSimian) / float64(countHuman)

	return protocols.StatsResponse {
		CountHumanDNA: countHuman,
		CountSimianDNA: countSimian,
		Ratio: ratio,
	}, nil
}
