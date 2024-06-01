package controllers

import (
	"GO/database"
	"GO/helpers"
	"GO/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := helpers.HashPassword(input.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    input.Password = hashedPassword
    input.CreatedAt = time.Now()
    input.UpdatedAt = time.Now()

    if err := database.DB.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": input})
}

func LoginUser(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
        return
    }

    if !helpers.CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password salah"})
        return
    }

    token, err := helpers.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
    userID, _ := c.Params.Get("userId")
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var user models.User
    if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User tidak dapat ditemukan"})
        return
    }

    if input.Password != "" {
        hashedPassword, err := helpers.HashPassword(input.Password)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        input.Password = hashedPassword
    }

    database.DB.Model(&user).Updates(input)
    c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
    userID, _ := c.Params.Get("userId")
    var user models.User
    if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User tidak dapat ditemukan"})
        return
    }

    database.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"data": "User dihapus"})
}
