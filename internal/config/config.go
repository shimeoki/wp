package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DB    DB
	Store Store
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

	db, store := "db", "store"

	v.SetDefault(db, DB{})
	v.SetDefault(store, Store{})

	log.Println(v.AllSettings())

	return &Config{
		DB:    getDB(v.Sub(db)),
		Store: getStore(v.Sub(store)),
	}
}

type DB struct {
	DataSourceName string
}

func getDB(v *viper.Viper) DB {
	dsn := "data-source-name"

	v.SetDefault(dsn, "db.sqlite")

	return DB{
		DataSourceName: v.GetString(dsn),
	}
}

type Store struct {
	Path string
}

func getStore(v *viper.Viper) Store {
	path := "path"

	v.SetDefault(path, "test/store")

	return Store{
		Path: v.GetString(path),
	}
}
