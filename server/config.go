package server

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type dsnConfig struct {
	host     string
	user     string
	password string
	db       string
	port     string
}

var (
	DriverName              string
	PosgresDBHost           string
	PosgresDBUser           string
	PosgresDBPassword       string
	PosgresDBName           string
	PosgresDBPort           string
	AppPort                 int
	DbMaxOpenConnection     int
	DbMaxIdleConnection     int
	DbConnectionMaxLifeTime time.Duration
)

func SecretConfig() {
	AppPort = viper.GetInt("APP_PORT")
	DriverName = viper.GetString("DRIVER_NAME")
	PosgresDBHost = viper.GetString("POSTGRES_HOST")
	PosgresDBUser = viper.GetString("POSTGRES_USER")
	PosgresDBPassword = viper.GetString("POSTGRES_PASSWORD")
	PosgresDBName = viper.GetString("POSTGRES_DB_NAME")
	PosgresDBPort = viper.GetString("POSTGRES_PORT")
	DbMaxOpenConnection = viper.GetInt("POSTGRES_DB_MAX_OPEN_CONNECTION")
	DbMaxIdleConnection = viper.GetInt("POSTGRES_DB_MAX_IDLE_CONNECTION")
}

func GetPostgresDSN() string {
	return writePostgreDSN(dsnConfig{
		host:     PosgresDBHost,
		user:     PosgresDBUser,
		password: PosgresDBPassword,
		db:       PosgresDBName,
		port:     PosgresDBPort,
	})
}

func writePostgreDSN(dsn dsnConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dsn.host, dsn.user, dsn.password, dsn.db, dsn.port)
}
