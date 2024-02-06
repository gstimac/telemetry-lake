package main

import (
	"go-study/config"
	"go-study/server"
)

func main() {
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	config.Init("test")
	server.Init()
}

func s3_json_events_put() {

}

func validate_hmac_secret() {

}
