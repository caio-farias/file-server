package handler

import (
	"file-server/src/db"
	"io"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

var bucketClient *storage.BucketHandle = db.BucketClient

func ServeFromGCS(c *gin.Context) {

	// Get the requested file name
	filename := c.Param("filename")

	// Check if the file exists
	obj := bucketClient.Object(filename)
	_, err := obj.Attrs(c.Request.Context())
	if err != nil {
		c.String(http.StatusNotFound, "File not found")
		return
	}

	// Set the headers for the response
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")

	// Serve the file from GCS
	reader, err := obj.NewReader(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer reader.Close()

	c.Stream(func(w io.Writer) bool {
		buf := make([]byte, 1024*1024)
		n, err := reader.Read(buf)
		if err != nil {
			return false
		}

		c.Writer.Write(buf[:n])
		return true
	})
}
