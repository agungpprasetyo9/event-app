package controllers

import (
	"net/http"

	"example.com/event-app/config"
	"example.com/event-app/models"
	"github.com/gin-gonic/gin"
)

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	config.DB.Create(&event)
	context.JSON(http.StatusCreated, gin.H{
		"message": "Data berhasil di buat",
		"event":   &event,
	})
}

func GetEvents(context *gin.Context) {
	var event []models.Event

	config.DB.Find(&event)
	context.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dibuat",
		"event":   event,
	})
}

func GetEventsbyId(context *gin.Context) {
	var event models.Event
	paramsId := context.Param("id")

	var eventData = config.DB.First(&event, paramsId).Error
	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event Tidak ditemukan",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Data tampil detail event",
		"event":   event,
	})
}

func UpdateEvent(context *gin.Context) {
	var event models.Event
	paramsId := context.Param("id")

	var eventData = config.DB.First(&event, paramsId).Error
	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event Tidak ditemukan",
		})
		return
	}

	var input models.Event
	err := context.ShouldBindJSON(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	config.DB.Model(&event).Updates(input)
	context.JSON(http.StatusOK, gin.H{
		"message": "Data Berhasil diupdate",
		"event":   event,
	})
}

func DeleteEvent(context *gin.Context) {
	var event models.Event
	paramsId := context.Param("id")

	var eventData = config.DB.First(&event, paramsId).Error
	if eventData != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Event Tidak ditemukan",
		})
		return
	}

	config.DB.Unscoped().Delete(&event)
	context.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil di delete",
	})
}
