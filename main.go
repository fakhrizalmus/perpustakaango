package main

import (
	"os"

	"github.com/fakhrizalmus/perpustakaango/initializers"
	model "github.com/fakhrizalmus/perpustakaango/models"
	"github.com/fakhrizalmus/perpustakaango/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	InitiateRoute()
}

func InitiateRoute() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
		port   = os.Getenv("PORT")
	)
	model.ConnectDatabase()
	routes.Initiate(api)

	router.Run(port)
}
