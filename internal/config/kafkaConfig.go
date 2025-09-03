package config

import "os"

type KafkaConfig struct {
	KafkaBroker string
	KafkaTopic  string
	DBConn      string
}

func Load() *KafkaConfig {
	return &KafkaConfig{
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		KafkaTopic:  os.Getenv("KAFKA_TOPIC"),
		DBConn:      os.Getenv("POSTGRES_DSN"),
	}
}
