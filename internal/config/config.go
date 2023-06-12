package config

import "sync"

var (
	once   sync.Once
	config *Config
)

type (
	Config struct {
		Transport TransportConfig `mapstructure:"mq"`
	}

	TransportConfig struct {
		HTTP  HTTPConfig  `mapstructure:"http"`
		GRPC  GRPCConfig  `mapstructure:"grpc"`
		Kafka KafkaConfig `mapstructure:"kafka"`
	}

	HTTPConfig struct{}

	GRPCConfig struct{}

	KafkaConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Topic    string `mapstructure:"topic"`
		MaxBytes int    `mapstructure:"maxBytes"`
	}
)

func Get() *Config {
	once.Do(func() {
		__init()
	})

	return config
}

func __init() {

}
