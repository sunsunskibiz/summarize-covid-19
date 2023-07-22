package service

import (
	"github.com/openlyinc/pointy"
	"github.com/sunsunskibiz/summarize-covid-19/model"
)

func CovidSummary(covids []model.Covid) model.CovidSummary {
	province := make(map[string]int)
	ageGroup := make(map[string]int)

	for _, cov := range covids {
		if prov := pointy.StringValue(cov.Province, ""); prov != "" {
			province[prov] += 1
		}

		if cov.Age == nil {
			ageGroup["N/A"] += 1
			continue
		}
		switch {
		case *cov.Age >= 0 && *cov.Age <= 30:
			ageGroup["0-30"] += 1
		case *cov.Age >= 31 && *cov.Age <= 60:
			ageGroup["31-60"] += 1
		case *cov.Age > 60:
			ageGroup["61+"] += 1
		}
	}

	return model.CovidSummary{Province: province, AgeGroup: ageGroup}
}
