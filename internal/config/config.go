package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DB    DB
	Store Store
}

type DB struct {
	DataSourceName string
}

type Store struct {
	Path string
}

func Load(file string) *Config {
	v := viper.New()

	if file != "" {
		v.SetConfigFile(file)
	}

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	dir, err := os.UserConfigDir()
	if err == nil {
		v.AddConfigPath(dir)
	}

	v.AddConfigPath(".")

	v.SetEnvPrefix("wp")
	v.AutomaticEnv()

	v.ReadInConfig()

	// code below is bad, but viper doesn't work with Sub() for some reason

	v.SetDefault("db.data-source-name", "db.sqlite")
	v.SetDefault("store.path", "test/store")

	return &Config{
		DB: DB{
			DataSourceName: v.GetString("db.data-source-name"),
		},

		Store: Store{
			Path: v.GetString("store.path"),
		},
	}
}
