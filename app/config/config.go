package config

import (
	"github.com/spf13/viper"
)

var c *viper.Viper

// Init : Reads configuration files.
func Init() error {
	c = viper.New()
	c.SetConfigName("dev")
	c.SetConfigType("yaml")
	c.AddConfigPath("./config/environments")
	c.AutomaticEnv()
	if err := c.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// GetConfig : Returns configuration values.
func GetConfig() *viper.Viper {
	return c
}
