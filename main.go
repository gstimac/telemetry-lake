package main

import (
	"go-study/config"
	"go-study/server"
)

func main() {
	config.Init("development")
	server.Init()
}

func s3_json_events_put() {

}

func validate_hmac_secret() {

}
