package model

type Covid struct {
	ConfirmDate    *string `json:"ConfirmDate"`
	No             *int    `json:"No"`
	Age            *int    `json:"Age"`
	Gender         *string `json:"Gender"`
	GenderEn       *string `json:"GenderEn"`
	Nation         *string `json:"Nation"`
	NationEn       *string `json:"NationEn"`
	Province       *string `json:"Province"`
	ProvinceId     *int    `json:"ProvinceId"`
	District       *string `json:"District"`
	ProvinceEn     *string `json:"ProvinceEn"`
	StatQuarantine *int    `json:"StatQuarantine"`
}

type CovidSummary struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}

type CovidResponse struct {
	Data []Covid `json:"Data"`
}