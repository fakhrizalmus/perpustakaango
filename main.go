package main

import (
	model "github.com/fakhrizalmus/perpustakaango/models"
	"github.com/fakhrizalmus/perpustakaango/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	InitiateRoute()
}

func InitiateRoute() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)
	model.ConnectDatabase()
	routes.Initiate(api)

	router.Run(":8080")
}
