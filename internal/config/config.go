package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
}

func Load() Config {
	viper.SetDefault("PORT", "8080")
	viper.AutomaticEnv()

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	return c
}
