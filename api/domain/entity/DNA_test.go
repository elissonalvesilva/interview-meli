package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDNA(t *testing.T) {
	t.Run("Should return a error if dna is empty", func(t *testing.T) {
		t.Parallel()

		var dna []string
		resp, err := NewDNA(dna)
		assert.Error(t, err, ErrEmptyDNA)
		assert.Nil(t, resp)
	})

	t.Run("Should return dna sequence", func(t *testing.T) {
		t.Parallel()
		dna := []string{"AAATTT", "AAATTT", "GTATA"}
		resp, err := NewDNA(dna)
		assert.Nil(t, err)
		assert.Equal(t, &DNASequence{DNA: dna}, resp)
	})
}
