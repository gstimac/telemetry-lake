package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../internal/config/")
	config.AddConfigPath("internal/config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("Error parsing configuration file!")
	}
}

func GetConfig() *viper.Viper {
	return config
}
