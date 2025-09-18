package config

type Config struct {
	DB *DB
}

type DB struct {
	DataSourceName string
}
