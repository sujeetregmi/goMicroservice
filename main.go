package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sujeetregmi/goMicroservice/api"
	"github.com/sujeetregmi/goMicroservice/database"
)

var router *gin.Engine

func init() {
	database.NewPostgreSQLClient()

}

func main() {
	r := api.SetUpRoutes()
	r.Run(":5000")
}
