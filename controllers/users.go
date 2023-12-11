package controllers

import (
	"fmt"
	"hrms/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexUsers(c *gin.Context) {
	var users []models.User
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		panic("String conversion Error: " + err.Error())
	}
	models.DB.Limit(limit).Offset(offset).Find(&users)
	c.JSON(200, users)
}

func StoreUsers(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind json: " + err.Error()})
		return
	}

	id := models.DB.Save(&user)
	//id := models.DB.Create(&user)
	if id == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to store to database"})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func SearchByTicketNo(c *gin.Context) {
	var users []models.User
	searchTerm := c.Param("searchTerm")

	models.DB.Limit(20).Offset(1).Preload("Document").Where("ticket_no LIKE ?", fmt.Sprintf("%%%s%%", searchTerm)).Find(&users)
	c.JSON(http.StatusOK, users)
}

func Login(c *gin.Context) {
	var user models.User
	type ReqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req ReqData

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password can't be empty"})
		return
	}

	result := models.DB.First(&user).Select("WHERE username = ? AND password = ?", req.Username, req.Password)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": result.Error})
			return
		}
	}

	c.JSON(http.StatusOK, user)
}
