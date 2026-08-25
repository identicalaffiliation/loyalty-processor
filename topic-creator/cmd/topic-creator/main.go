package main

import (
	"flag"
	"log"
	"os"

	"github.com/identicalaffiliation/loyalty-processor/topic-creator/internal/config"
	"github.com/identicalaffiliation/loyalty-processor/topic-creator/internal/creator"
	"github.com/identicalaffiliation/loyalty-processor/topic-creator/pkg/logger"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	slogger, err := logger.NewLogger(&cfg.LoggerConfig)
	if err != nil {
		log.Fatal(err)
	}

	if err := creator.CreateKafkaTopics(slogger, &cfg.KafkaConfig); err != nil {
		slogger.Error(
			"error", err,
		)
		os.Exit(1)
	}

	slogger.Debug("kafka topics was created successfully!")
}
