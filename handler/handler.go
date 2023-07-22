package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CovidSummary(c *gin.Context) {
	// TODO: Get data from API

	// TODO: Logic Summary

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
