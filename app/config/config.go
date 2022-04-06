package config

import (
	"github.com/spf13/viper"
)

var c *viper.Viper

// Init : Reads configuration files.
func Init(env string) {
	c = viper.New()
	c.SetConfigFile("yaml")
	c.SetConfigName(env)
	c.AddConfigPath("config/environments/")
	if err := c.ReadInConfig(); err != nil {
		panic(err)
	}
}

// GetConfig : Returns configuration values.
func GetConfig() *viper.Viper {
	return c
}
