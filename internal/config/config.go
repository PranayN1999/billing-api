package config

import "github.com/spf13/viper"

// Config holds all runtime settings read from environment variables (or .env).
type Config struct {
	Port      string `mapstructure:"PORT"`
	DBURL     string `mapstructure:"DATABASE_URL"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

// Load reads env / .env and fills Config.
// Fatal-panics on any unmarshalling error so we fail fast on bad config.
func Load() Config {
	// defaults
	viper.SetDefault("PORT", "8080")

	// Optional .env file in project root
	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig() // ignore error—file is optional

	viper.AutomaticEnv() // override with real env-vars

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	return c
}
