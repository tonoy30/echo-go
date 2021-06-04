package config

import (
	"log"

	"github.com/spf13/viper"
)

type Settings struct {
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBDriver    string `mapstructure:"DB_DRIVER"`
	DBPassword  string `mapstructure:"DB_PASS"`
	DBUser      string `mapstructure:"DB_USER"`
	DBName      string `mapstructure:"DB_NAME"`
	Environment string `mapstructure:"ENV"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	JWTExpires  string `mapstructure:"JWT_EXPIRES"`
}

func NewSettings() *Settings {
	var settings Settings

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Println("No env file found", err)
	}
	err = viper.Unmarshal(&settings)

	if err != nil {
		log.Println("Error: while trying to unmarshal configuration", err)
	}
	return &settings
}
