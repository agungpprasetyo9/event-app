package controllers

import (
	"net/http"
	"os"
	"time"

	"example.com/event-app/config"
	"example.com/event-app/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthInputRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type AuthInputLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func RegisterUser(context *gin.Context) {
	var input AuthInputRegister

	err := context.ShouldBindJSON(&input)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if errHash != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//simpan ke database
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	userCreated := config.DB.Create(&user).Error

	if userCreated != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "email mungsinn sudah terdaftar",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil register",
		"user": gin.H{
			"id":     user.ID,
			"name":   user.Name,
			"email":  user.Email,
			"events": user.Events,
		},
	})
	return
}

func LoginUser(context *gin.Context) {
	var input AuthInputLogin

	// validation

	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	var user models.User
	userData := config.DB.Where("email = ?", input.Email).First(&user).Error
	if userData != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "email belum terdaftar",
		})
	}

	errMatchPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if errMatchPassword != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "password salah",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "gagal membuat token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "login berhasil",
		"token":   tokenString,
		"user": gin.H{
			"id":     user.ID,
			"name":   user.Name,
			"email":  user.Email,
			"events": user.Events,
		},
	})
}
