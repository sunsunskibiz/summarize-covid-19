package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnMarshalCovid(t *testing.T) {
	covidJson := []byte(`{
    "ConfirmDate": "2021-05-03",
    "No": null,
    "Age": 24,
    "Gender": null,
    "GenderEn": null,
    "Nation": null,
    "NationEn": "India",
    "Province": "Sisaket",
    "ProvinceId": 62,
    "District": null,
    "ProvinceEn": "Sisaket",
    "StatQuarantine": 15
}
`)

	var covid Covid
	err := json.Unmarshal(covidJson, &covid)

	assert.Nil(t, err)
	assert.Equal(t, 24, *covid.Age)
	assert.Equal(t, "India", *covid.NationEn)
}

func TestMarsholCovidSummary(t *testing.T) {
	covSum := CovidSummary{
		Province: map[string]int{
			"Samut Sakhon": 3613,
			"Bangkok":      2774,
		},
		AgeGroup: map[string]int{
			"0-30": 300,
			"N/A":  4,
		},
	}
	covSumJson, err := json.Marshal(covSum)
	assert.Nil(t, err)

	expectCovSumJson := `{"Province":{"Bangkok":2774,"Samut Sakhon":3613},"AgeGroup":{"0-30":300,"N/A":4}}`
	assert.Equal(t, expectCovSumJson, string(covSumJson))
}
