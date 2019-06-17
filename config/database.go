package config

import (
	"database/sql"
	"github.com/spf13/viper"
)

type Database struct {
	Host              string
	User              string
	Password          string
	DBName            string
	Port              int
	ReconnectRetry    int
	ReconnectInterval int64
	DebugMode         bool
}

// LoadDBConfig load database configuration
func LoadDBConfig(name string) Database {
	db := viper.Sub("database." + name)
	conf := Database{
		Host:              db.GetString("host"),
		User:              db.GetString("user"),
		Password:          db.GetString("password"),
		DBName:            db.GetString("db_name"),
		Port:              db.GetInt("port"),
		ReconnectRetry:    db.GetInt("reconnect_retry"),
		ReconnectInterval: db.GetInt64("reconnect_interval"),
		DebugMode:         db.GetBool("debug"),
	}
	return conf
}

func SqliteOpen() (*sql.DB, error) {
	conf := LoadDBConfig("sqlite")
	db, err := sql.Open("sqlite3", AppPath+"/"+conf.DBName+".db")
	return db, err
}
