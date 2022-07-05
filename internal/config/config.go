package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env:"ST_BOT_IS_DEBUG" env-default:"false"`
	Listen  struct {
		Type string `yaml:"type" env:"ST_APP_TYPE" env-required:"port"`
		Port string `yaml:"port" env:"ST_APP_PORT" env-required:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host        string `yaml:"host" env:"ST_BOT_RABBIT_HOST" env-required:"true"`
	Port        string `yaml:"port" env:"ST_BOT_RABBIT_PORT" env-required:"true"`
	Database    string `yaml:"database" env:"ST_APP_DATABASE" env-required:"true"`
	Username    string `yaml:"username" env:"ST_BOT_RABBIT_USERNAME" env-required:"true"`
	Password    string `yaml:"password" env:"ST_BOT_RABBIT_PASSWORD" env-required:"true"`
	MaxAttempts int    `yaml:"maxAttempts" env:"ST_BOT_RABBIT_PASSWORD" env-required:"5"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
