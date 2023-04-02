package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var StoragePath = os.Getenv("STORAGE_PATH")

func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	filepath := fmt.Sprintf("./", StoragePath, filename)

	// Check if the file exists
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		c.String(http.StatusNotFound, "File not found")
		return
	}

	// Set the headers for the response
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filepath)
}

func UploadFileHandler(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}

	c.String(http.StatusOK, "Uploaded successfully %d files with fields name=%s and email=%s.", len(files), name, email)
}
