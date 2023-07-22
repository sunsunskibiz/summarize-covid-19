package service

import (
	"testing"

	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
	"github.com/sunsunskibiz/summarize-covid-19/model"
)

func TestCovidSummary(t *testing.T) {
	covs := []model.Covid{
		{
			Province: pointy.String("Samut Sakhon"),
			Age:      pointy.Int(32),
		},
		{
			Province: pointy.String("Bangkok"),
			Age:      pointy.Int(44),
		},
		{
			Province: pointy.String("Samut Sakhon"),
			Age:      pointy.Int(12),
		},
		{
			Province: pointy.String("Bangkok"),
			Age:      pointy.Int(98),
		},
		{
			Province: pointy.String("Bangkok"),
			Age:      pointy.Int(23),
		},
		{
			Province: pointy.String(""),
			Age:      pointy.Int(35),
		},
		{
			Age: pointy.Int(72),
		},
		{
			Age: pointy.Int(0),
		},
		{
			Province: pointy.String("Bangkok"),
		},
	}

	covSum := CovidSummary(covs)

	assert.Equal(t, 4, covSum.Province["Bangkok"])
	assert.Equal(t, 2, covSum.Province["Samut Sakhon"])
	assert.Equal(t, 3, covSum.AgeGroup["0-30"])
	assert.Equal(t, 3, covSum.AgeGroup["31-60"])
	assert.Equal(t, 2, covSum.AgeGroup["61+"])
	assert.Equal(t, 1, covSum.AgeGroup["N/A"])
}
