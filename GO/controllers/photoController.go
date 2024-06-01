package controllers

import (
	"GO/database"
	"GO/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
    var input models.Photo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, _ := c.Get("userID")
    input.UserID = userID.(uint)
    input.CreatedAt = time.Now()
    input.UpdatedAt = time.Now()

    if err := database.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": input})
}

func GetPhotos(c *gin.Context) {
    var photos []models.Photo
    database.DB.Find(&photos)
    c.JSON(http.StatusOK, gin.H{"data": photos})
}

func UpdatePhoto(c *gin.Context) {
    photoID, _ := c.Params.Get("photoId")
    var input models.Photo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var photo models.Photo
    if err := database.DB.Where("id = ?", photoID).First(&photo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
        return
    }

    userID, _ := c.Get("userID")
    if photo.UserID != userID {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
        return
    }

    database.DB.Model(&photo).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": photo})
}

func DeletePhoto(c *gin.Context) {
    photoID, _ := c.Params.Get("photoId")
    var photo models.Photo
    if err := database.DB.Where("id = ?", photoID).First(&photo).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
        return
    }

    userID, _ := c.Get("userID")
    if photo.UserID != userID {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
        return
    }

    database.DB.Delete(&photo)
    c.JSON(http.StatusOK, gin.H{"data": "Photo deleted"})
}

