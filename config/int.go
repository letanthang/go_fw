package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Schema struct {
	Mongo struct {
		Host     string
		Username string
		Password string
	}

	Encryption struct {
		JWTSecret string
	}
}

var Config Schema

func init() {
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
