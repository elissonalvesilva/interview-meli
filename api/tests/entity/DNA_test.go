package entity

import (
	"github.com/elissonalvesilva/interview-meli/api/domain/entity"
	"reflect"
	"testing"
)

func TestNewDNA(t *testing.T) {
	type args struct {
		dna []string
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.DNASequence
		wantErr bool
	}{
		{
			"Should return a error if dna is empty",
			args{},
			nil,
			true,
		},
		{
			"Should return dna sequence",
			args{dna: []string{"AAATTT", "AAATTT", "GTATA"}},
			&entity.DNASequence{DNA: []string{"AAATTT", "AAATTT", "GTATA"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := entity.NewDNA(tt.args.dna)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDNA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDNA() got = %v, want %v", got, tt.want)
			}
		})
	}
}
