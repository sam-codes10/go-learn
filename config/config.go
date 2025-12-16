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

func Start(env string) Config {

	v := GetConfig()

	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath("./resources")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg EnvConfig

	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode config into struct: %v", err)
	}

	var config Config
	switch env {
	case "local":
		config.Kafka = cfg.Local.Kafka
		config.SMTP = cfg.Local.SMTP
	default:
		log.Panic("unknown env: ", env)
	}

	return config
}

type KafkaConfig struct {
	KafkaBrokers []string `mapstructure:"brokers"`
}

type SMTPConfig struct {
	SMTPHost string `mapstructure:"host"`
	SMTPPort int    `mapstructure:"port"`
	SMTPUser string `mapstructure:"user"`
	SMTPPass string `mapstructure:"pass"`
}

type Config struct {
	Kafka KafkaConfig `mapstructure:"kafka"`
	SMTP  SMTPConfig  `mapstructure:"smtp"`
}

type EnvConfig struct {
	Local Config `mapstructure:"local"`
}
