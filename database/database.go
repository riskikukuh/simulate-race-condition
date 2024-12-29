package database

import (
	"fmt"
	"log/slog"
	"simulation-race-condition/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type WrapDB struct {
	PostgreDB *gorm.DB
}

func InitDatabase(env *config.EnvironmentVariable) (*WrapDB, error) {
	slog.Info("Connecting to PostgreSql")
	postgreSqlConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		env.Database.PG_HOST,
		env.Database.PG_USERNAME,
		env.Database.PG_PASSWORD,
		env.Database.PG_NAME,
		env.Database.PG_PORT,
		env.Database.PG_TIMEZONE,
	)
	postgreDB, err := gorm.Open(postgres.Open(postgreSqlConnection), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	sqlDB, err := postgreDB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(11)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(12)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		slog.Info("Failed connect to PostgreSql")
		return nil, err
	}
	slog.Info("Database PostgreSql is connected")

	return &WrapDB{
		PostgreDB: postgreDB,
	}, nil
}
