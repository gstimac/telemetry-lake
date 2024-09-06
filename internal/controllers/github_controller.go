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
	"telemetry-lake/internal/config"
	"telemetry-lake/internal/models"
	"time"
)

type GithubEvent models.GithubEvent
type GithubController struct{}

func CreateGithubEvent(reader io.Reader) GithubEvent {
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

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	ghevent := GithubEvent{
		ID:          rng.Int(),
		Name:        "macka",
		Age:         14,
		Description: "debela macka",
	}

	objectName := "githubevent" + strconv.Itoa(rng.Int()) + ".jsonl"
	contentType := "application/json"

	// Upload the test file with FPutObject. With size -1 this will use memory, fix this later
	info, err := mc.PutObject(ctx, bucket, objectName, reader, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return ghevent
}

func (h GithubController) WriteEvent(c *gin.Context) {
	CreateGithubEvent(c.Request.Body)
	c.JSON(200, gin.H{
		"message": "written",
	})
}

func (h GithubController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
