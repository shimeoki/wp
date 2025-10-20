package config

import (
	"os"
	"path"

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

	dir := "."

	configs, err := os.UserConfigDir()
	if err == nil {
		dir = path.Join(configs, "wp")
	}

	if err := os.MkdirAll(dir, 0700); err != nil {
		dir = "."
	}

	v.AddConfigPath(dir)

	v.SetEnvPrefix("wp")
	v.AutomaticEnv()

	v.ReadInConfig()

	v.SetDefault("db.data-source-name", path.Join(dir, "db.sqlite"))
	v.SetDefault("store.path", path.Join(dir, "store"))

	return &Config{
		DB: DB{
			DataSourceName: v.GetString("db.data-source-name"),
		},

		Store: Store{
			Path: v.GetString("store.path"),
		},
	}
}
