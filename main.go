package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subbarao/transformer/pkg/server"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	port := viper.GetInt("PORT")
	quickPlayURL := viper.GetString("QUICK_PLAY_URL")
	server.Start(port, quickPlayURL)
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
