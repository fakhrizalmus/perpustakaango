package routes

import (
	"github.com/fakhrizalmus/perpustakaango/controllers/authcontroller"
	"github.com/gin-gonic/gin"
)

func RouteAuth(r *gin.RouterGroup) {
	r.POST("/register", authcontroller.Register)
	r.POST("/login", authcontroller.Login)
	r.GET("/logout", authcontroller.Logout)
}
