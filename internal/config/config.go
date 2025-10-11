package config

type Config struct {
	DB
}

type DB struct {
	DataSourceName string
}
