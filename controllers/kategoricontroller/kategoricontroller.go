package kategoricontroller

import (
	"net/http"

	"github.com/fakhrizalmus/perpustakaango/common"
	model "github.com/fakhrizalmus/perpustakaango/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	var kategori []model.Kategori
	model.DB.Find(&kategori)
	c.JSON(http.StatusOK, common.APIResponse{
		Status: true,
		Data:   kategori,
	})
}

func GetByID(c *gin.Context) {
	var kategori model.Kategori
	id := c.Param("id")
	err := model.DB.First(&kategori, id).Error
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
				Status:  false,
				Message: "Data Not Found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
				Status:  false,
				Message: err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Status: true,
		Data:   kategori,
	})
}

func Create(c *gin.Context) {
	var kategori model.Kategori
	err := c.ShouldBindBodyWithJSON(&kategori)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}
	model.DB.Create(&kategori)
	c.JSON(http.StatusOK, common.APIResponse{
		Status: true,
		Data:   kategori,
	})
}

func Edit(c *gin.Context) {
	var kategori model.Kategori
	id := c.Param("id")
	if model.DB.Model(&kategori).Where("id = ?", id).Updates(&kategori).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Status:  false,
			Message: "ID Not Found"})
		return
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Status: true,
		Data:   kategori,
	})
}

func Delete(c *gin.Context) {
	var kategori model.Kategori
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Status:  false,
			Message: "ID tidak valid"})
		return
	}

	if model.DB.First(&kategori, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
			Status:  false,
			Message: "Produk tidak ditemukan"})
		return
	}

	if model.DB.Delete(&kategori).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Status:  false,
			Message: "Gagal Hapus"})
		return
	}

	c.JSON(http.StatusOK, common.APIResponse{
		Status:  true,
		Message: "Delete success",
	})
}
