package main

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/subbarao/transformer/pkg/server"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	quickPlayURL := "https://data-store-cdn.cms-stag.amdvids.com"
	app := server.New(quickPlayURL)
	n, _ := strconv.Atoi(port)
	app.Run(n)
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}
