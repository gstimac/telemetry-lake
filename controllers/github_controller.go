package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"math/rand"
	"strconv"
	"telemetry-lake/config"
	"time"
)

type GithubController struct{}

func (h GithubController) WriteEvent(c *gin.Context) {
	putBlob(c.Request.Body)
	c.JSON(200, gin.H{
		"message": "written",
	})
}

func (h GithubController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func putBlob(ioReader io.Reader) {
	cfg := config.GetConfig()
	ctx := context.Background()
	var endpoint = cfg.GetString("minio.endpoint")
	var accessKeyID = cfg.GetString("minio.user.access-key-id")
	var secretAccessKey = cfg.GetString("minio.user.secret-access-key")
	var bucket = cfg.GetString("minio.bucket")
	var useSSL = cfg.GetBool("minio.useSSL")

	mc, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	rand.Seed(time.Now().UnixNano())
	objectName := "testdata" + strconv.Itoa(rand.Int())
	contentType := "application/octet-stream"

	// Upload the test file with FPutObject. With size -1 this will use memory, fix this later
	info, err := mc.PutObject(ctx, bucket, objectName, ioReader, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}
