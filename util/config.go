package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	SecretKey string `mapstructure:"SECRET_KEY"`
}
func LoadConfig(path string) (config Config, err error){
	// add path to get the env's if not set
	viper.AddConfigPath(path)
	// specify file details, the name and the extension of the file
	viper.SetConfigName("appEnv")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		return 
	}
	err = viper.Unmarshal(&config)
	return
}