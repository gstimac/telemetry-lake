package main

import (
	"telemetry-lake/internal/config"
	"telemetry-lake/internal/server"
)

func main() {
	config.Init("development")
	server.Init()
}

func s3_json_events_put() {

}

func validate_hmac_secret() {

}
