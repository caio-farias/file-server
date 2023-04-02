package service

import (
	"bytes"
	"file-server/src/config"
	"file-server/src/db"
	"io/ioutil"

	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	bucketClient = db.BucketClient
	bucketName   = os.Getenv("GCP_BUCKET_NAME")
)

// createFile creates a file in Google Cloud Storage.
func CreateFile(fileName string, c *gin.Context) {
	config.LOGGER.Info("Creating ", zap.String("file\n", fileName))

	wc := bucketClient.Object(fileName).NewWriter(c.Request.Context())
	wc.ContentType = "text/plain"
	wc.Metadata = map[string]string{
		"x-goog-meta-foo": "foo",
		"x-goog-meta-bar": "bar",
	}

	if _, err := wc.Write([]byte("abcde\n")); err != nil {
		config.LOGGER.Error("createFile: unable to write data to bucket ", zap.String("bucket ", bucketName), zap.String("file: ", fileName), zap.Error(err))
		return
	}
	if _, err := wc.Write([]byte(strings.Repeat("f", 1024*4) + "\n")); err != nil {
		config.LOGGER.Error("createFile: unable to write data to bucket ", zap.String("bucket ", bucketName), zap.String("file: ", fileName), zap.Error(err))
		return
	}
	if err := wc.Close(); err != nil {
		config.LOGGER.Error("createFile: unable to close, ", zap.String("bucket ", bucketName), zap.String("file: ", fileName), zap.Error(err))
		return
	}
}

// readFile reads the named file in Google Cloud Storage.
func ReadFile(fileName string, c *gin.Context) {
	config.LOGGER.Info("\nAbbreviated file content (first line and last 1K):\n")

	rc, err := bucketClient.Object(fileName).NewReader(c.Request.Context())
	if err != nil {
		config.LOGGER.Error("readFile: unable to open file from ", zap.String("bucket ", bucketName), zap.String("file: ", fileName), zap.Error(err))
		return
	}
	defer rc.Close()
	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		config.LOGGER.Error("readFile: unable to read data from bucket ", zap.String("bucket ", bucketName), zap.String("file: ", fileName), zap.Error(err))
		return
	}

	config.LOGGER.Info("\n", zap.ByteString("", bytes.SplitN(slurp, []byte("\n"), 2)[0]))
	if len(slurp) > 1024 {
		config.LOGGER.Info("...\n", zap.ByteString("", slurp[len(slurp)-1024:]))
	} else {
		config.LOGGER.Info("\n", zap.ByteString("", slurp))
	}
}

// deleteFiles deletes all the temporary files from a bucket created by this demo.
func DeleteFiles(cleanUp []string, c *gin.Context) {
	config.LOGGER.Info("\nDeleting files...\n")
	for _, file := range cleanUp {
		config.LOGGER.Info("\nDeleting \n", zap.String("file...", file))
		if err := bucketClient.Object(file).Delete(c.Request.Context()); err != nil {
			config.LOGGER.Error("deleteFiles: unable to delete bucket ", zap.String("bucket ", bucketName), zap.String("file: ", file), zap.Error(err))
			return
		}
	}
}
