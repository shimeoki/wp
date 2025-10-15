package config

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
