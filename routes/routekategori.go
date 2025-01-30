package routes

import (
	"github.com/fakhrizalmus/perpustakaango/controllers/kategoricontroller"
	"github.com/fakhrizalmus/perpustakaango/middlewares"
	"github.com/gin-gonic/gin"
)

func RouteKategori(r *gin.RouterGroup) {
	r.Use(middlewares.AuthMiddleware)
	r.GET("/", kategoricontroller.GetAll)
	r.GET("/:id", kategoricontroller.GetByID)
	r.POST("/", kategoricontroller.Create)
	r.PATCH("/:id", kategoricontroller.Edit)
	r.DELETE("/:id", kategoricontroller.Delete)
}
