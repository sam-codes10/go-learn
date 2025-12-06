package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type ViperConfig struct {
	viper *viper.Viper
}

var configInstance *ViperConfig
var singleton sync.Once

func GetConfig() *viper.Viper {
	singleton.Do(func() {
		configInstance = &ViperConfig{viper.New()}
	})
	return configInstance.viper
}

func Start(env string) EnvConfig {

	v := GetConfig()

	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./resources")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode config into struct: %v", err)
	}

	var config EnvConfig
	switch env {
	case "local":
		config.Kafka = cfg.Local.Kafka
		config.SMTP = cfg.Local.SMTP
	case "dev":
		config.Kafka = cfg.Dev.Kafka
		config.SMTP = cfg.Dev.SMTP
	case "prod":
		config.Kafka = cfg.Prod.Kafka
		config.SMTP = cfg.Prod.SMTP
	default:
		log.Panic("unknown env: ", env)
	}

	return config
}

type KafkaConfig struct {
	KafkaBrokers []string `mapstructure:"brokers"`
	KafkaTopic   string   `mapstructure:"topic"`
	KafkaGroup   string   `mapstructure:"group"`
}

type SMTPConfig struct {
	SMTPHost string `mapstructure:"host"`
	SMTPPort int    `mapstructure:"port"`
	SMTPUser string `mapstructure:"user"`
	SMTPPass string `mapstructure:"pass"`
}

type EnvConfig struct {
	Kafka KafkaConfig `mapstructure:"kafka"`
	SMTP  SMTPConfig  `mapstructure:"smtp"`
}

type Config struct {
	Local EnvConfig `mapstructure:"local"`
	Dev   EnvConfig `mapstructure:"dev"`
	Prod  EnvConfig `mapstructure:"prod"`
}
