package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	v *viper.Viper

	DB struct {
		DataSourceName string
	}

	Store struct {
		Path string
	}

	Log struct {
		Path string
	}
}

func New() *Config {
	v := viper.New()

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

	v.SetDefault("db.data-source-name", path.Join(dir, "db.sqlite"))
	v.SetDefault("store.path", path.Join(dir, "store"))
	v.SetDefault("log.path", path.Join(dir, "wp.log"))

	cfg := &Config{v: v}
	cfg.load() // defaults with no filesystem access

	return cfg
}

func (c *Config) load() {
	c.DB.DataSourceName = c.v.GetString("db.data-source-name")
	c.Store.Path = c.v.GetString("store.path")
	c.Log.Path = c.v.GetString("log.path")
}

func (c *Config) Load(file string) error {
	c.v.SetConfigFile(file)

	if err := c.v.ReadInConfig(); err != nil {
		return err
	}

	c.load()
	return nil
}
