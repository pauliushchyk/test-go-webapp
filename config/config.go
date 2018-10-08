package config

import (
	"os"

	"github.com/spf13/viper"
)

// Get returns configuration value by key
func Get(k string) string {
	viper.SetConfigName(getENV())
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if e := viper.ReadInConfig(); e != nil {
		panic("failed to read configuration value by key")
	}

	return viper.GetString(k)
}

func getENV() string {
	if env := os.Getenv("GO_ENV"); len(env) != 0 {
		return env
	}

	return "development"
}
