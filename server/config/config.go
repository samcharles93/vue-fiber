package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Supabase SupabaseConfig `mapstructure:"supabase"`
}

type ServerConfig struct {
	Port   string `mapstructure:"port"`
	Domain string `mapstructure:"domain"`
}

type SupabaseConfig struct {
	URL string `mapstructure:"url"`
	Key string `mapstructure:"key"`
}

var Cfg Config

func InitConfig() (*Config, error) {
	viper.SetConfigFile("../.local/config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read in config file: %v", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return &Cfg, nil
}
