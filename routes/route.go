package routes

import "github.com/gin-gonic/gin"

func Initiate(api *gin.RouterGroup) {
	kategori := api.Group("/kategoris")
	RouteKategori(kategori)
}
