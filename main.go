package main

import (
	"telemetry-lake/internal/config"
	"telemetry-lake/internal/server"
)

func main() {
	config.Init("development")
	server.Init()
}
