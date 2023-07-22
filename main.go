package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sunsunskibiz/summarize-covid-19/handler"
)

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/covid/summary", handler.CovidSummary)

	return r
}
