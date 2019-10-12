package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Schema struct {
	Mongo struct {
		Host     string `mapstructure:"Host"`
		Username string `mapstructure:"User"`
		Password string `mapstructure:"Password"`
	} `mapstructure:"MongoDB"`

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
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v\n", err))
	}
}
