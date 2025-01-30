package authcontroller

import (
	"net/http"
	"os"
	"time"

	"github.com/fakhrizalmus/perpustakaango/common"
	model "github.com/fakhrizalmus/perpustakaango/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var (
		req  model.User
		user model.User
	)

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	//cek email
	if err := model.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
				Status:  false,
				Message: "Email salah",
				Error:   err.Error(),
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
				Status:  false,
				Message: err.Error(),
			})
			return
		}
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
			Status:  false,
			Message: "Password salah",
			Error:   err.Error(),
		})
		return
	}

	// proses jwt
	expTime := time.Now().Add(time.Hour)

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": jwt.NewNumericDate(expTime),
	})
	//signed token
	tokenString, err := tokenAlgo.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}
	c.SetCookie("Authorization", tokenString, 3600, "/", "", false, true)

	c.JSON(200, gin.H{
		"message": "Cookie set successfully",
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{
		"message": "Logout successfully",
	})
}

func Register(c *gin.Context) {
	var (
		req model.User
	)

	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashPassword)

	if err := model.DB.Create(&req).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.APIResponse{
			Status:  false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Status: true,
		Data:   req,
	})
}
