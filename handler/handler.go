package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunsunskibiz/summarize-covid-19/model"
)

func CovidSummary(c *gin.Context) {
	covids, err := getCovids()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot get or process data from 'https://static.wongnai.com/devinterview/covid-cases.json'",
		})
		return
	}

	fmt.Printf("%+v\n", covids[0])

	// TODO: Logic Summary

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "pong",
	// })
}

func getCovids() ([]model.Covid, error) {
	resp, err := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	covidRes := model.CovidResponse{}
	err = json.Unmarshal(body, &covidRes)
	if err != nil {
		return nil, err
	}

	return covidRes.Data, nil
}
