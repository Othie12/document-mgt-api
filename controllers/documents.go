package controllers

import (
	"fmt"
	"hrms/models"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexDocuments(c *gin.Context) {
	var documents []models.Document
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		panic("String convesion error" + err.Error())
	}
	models.DB.Limit(limit).Offset(offset).Find(&documents)
	c.JSON(200, documents)
}

func StoreDocuments(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	application, _ := c.FormFile("applciation")
	certificate, _ := c.FormFile("certificate")
	appointment, _ := c.FormFile("appointment")
	discipline, _ := c.FormFile("discipline")
	others, _ := c.FormFile("others")

	dst := "./public"

	err := c.SaveUploadedFile(application, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save application: " + err.Error()})
		return
	}

	err = c.SaveUploadedFile(certificate, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save certificate: " + err.Error()})
		return
	}

	err = c.SaveUploadedFile(appointment, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save appointment: " + err.Error()})
		return
	}

	err = c.SaveUploadedFile(discipline, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save discipline: " + err.Error()})
		return
	}

	err = c.SaveUploadedFile(others, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save others: " + err.Error()})
		return
	}

	document := models.Document{
		UserID:      uint(id),
		Application: filepath.Join(dst, application.Filename),
		Certificate: filepath.Join(dst, certificate.Filename),
		Appointment: filepath.Join(dst, appointment.Filename),
		Discipline:  filepath.Join(dst, discipline.Filename),
		Others:      filepath.Join(dst, others.Filename),
	}

	docid := models.DB.Create(&document)
	if docid == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to save data" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, document)
}

func UpdateDocument(c *gin.Context) {
	var document models.Document

	//first convert the request's id param from string to uint64
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	document.ID = uint(id) //convert id into uint from uint64 and assign it to the doc

	//get the filename (application/certificate/othier)
	name := c.Param("name")

	dst := "./public"
	file, _ := c.FormFile("file")

	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to save to db: " + err.Error()})
		return
	}

	models.DB.Model(&document).Update(name, filepath.Join(dst, file.Filename))

	c.JSON(http.StatusOK, document)
}

func FetchByUserId(c *gin.Context) {
	var document models.Document

	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "failed to get user id: " + err.Error()})
		return
	}

	result := models.DB.Where("user_id = ?", userId).First(&document)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Failed to get record from database: " + result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, document)
}

func AddDocument(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	name := c.Param("name")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get user id: " + err.Error()})
		return
	}

	dst := "./public/"
	fmt.Println(dst)

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get the file: " + err.Error()})
		return
	}

	// Check for the existence of the file before saving
	filePath := filepath.Join(dst, file.Filename)
	if _, err := os.Stat(filePath); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "file with the same name already exists"})
		return
	}

	// Save the file
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save the file: " + err.Error()})
		return
	}

	document := models.Doc{
		UserID:   uint(userId),
		Name:     name,
		Filepath: filePath,
	}

	models.DB.Create(&document)
	c.JSON(http.StatusOK, document)
}
