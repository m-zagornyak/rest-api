package config

import (
	"fmt"
	"github.com/spf13/viper"
)

/*
type Config struct {
	Port    string `yaml:"port" env:"PORT" env-default:"8080"`
	IsDebug *bool
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
*/

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigFile("config")
	return viper.ReadInConfig() // Find and read the config file

}
func GetConfig() {
	if err := initConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
