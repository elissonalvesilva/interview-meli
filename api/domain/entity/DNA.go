package entity

import (
	"errors"
)

type DNASequence struct {
	DNA []string `json:"dna"`
}

const (
	errEmptyDNAMessage = "DNA must be pass"
)

var (
	ErrEmptyDNA = errors.New(errEmptyDNAMessage)
)

func NewDNA(dna []string) (*DNASequence, error) {

	if len(dna) <= 0 {
		return nil, ErrEmptyDNA
	}

	return &DNASequence{dna}, nil
}
