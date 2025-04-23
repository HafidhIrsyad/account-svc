package server

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func DBConnection() (*gorm.DB, error) {
	return makeConnection()
}

func makeConnection() (*gorm.DB, error) {
	dsn := GetPostgresDSN()

	dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("cannot gorm.Open on makeConnection")
		panic(err)
	}

	// Set connection pool settings
	sqlDB, err := dbGorm.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get DB instance: %w", err)
	}

	sqlDB.SetMaxOpenConns(DbMaxOpenConnection)
	sqlDB.SetMaxIdleConns(DbMaxIdleConnection)
	sqlDB.SetConnMaxLifetime(DbConnectionMaxLifeTime)

	return dbGorm, nil
}
