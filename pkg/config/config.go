package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() { 
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() 
	if err != nil { 
		panic(fmt.Errorf("fatal error config file: %w", err))   
	}
}