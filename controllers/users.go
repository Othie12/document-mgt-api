package controllers

import (
	"fmt"
	"hrms/models"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	models.DB.Limit(limit).Offset(offset).Order("username").Find(&users)
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

func ShowUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var user models.User
	user.ID = uint(id)

	result := models.DB.Preload("Docs").Find(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to get record from database: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchByTicketNo(c *gin.Context) {
	var users []models.User
	searchTerm := c.Param("searchTerm")

	models.DB.Limit(10).Preload("Docs").Where("ticket_no LIKE ?", fmt.Sprintf("%%%s%%", searchTerm)).Find(&users)
	c.JSON(http.StatusOK, users)
}

func SearchByName(c *gin.Context) {
	var users []models.User
	searchTerm := c.Param("searchTerm")

	models.DB.Limit(10).Preload("Docs").Where("username LIKE ?", fmt.Sprintf("%%%s%%", searchTerm)).Find(&users)
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

	result := models.DB.Where("username = ? AND password = ?", req.Username, req.Password).First(&user)
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

func UploadPhoto(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var user models.User

	result := models.DB.First(&user, uint(id))
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNoContent, gin.H{"error": "User not found"})
			log.Println(result.Error)
			return
		} else {
			c.JSON(http.StatusNoContent, gin.H{"error": result.Error})
			log.Println(result.Error)
			return
		}
	}

	dst := "./public/"

	// Check if the "uploads" directory exists, create it if not
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err = os.MkdirAll(dst, 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create 'uploads' directory"})
			return
		}
	}

	file, err := c.FormFile("file")
	if err != nil {
		panic("Error parsing formfile: " + err.Error())
	}

	file.Filename = fmt.Sprintf("%s_%s_%s", user.TicketNo, "pass", time.Now())

	filePath := filepath.Join(dst, file.Filename+".pdf")

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save to filesystem: " + err.Error()})
		log.Println(file.Filename)
		return
	}

	models.DB.Model(&user).Update("photo", fmt.Sprintf("/%s", filePath))

	c.JSON(http.StatusOK, user)
}
