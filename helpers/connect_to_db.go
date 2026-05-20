package helpers

import (
	"inapp_payment_service/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GormDB *gorm.DB

func ConnectToGorm() error {
	if GormDB == nil {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.PostgresHost(),
			config.PostgresUser(),
			config.PostgresPassword(),
			config.PostgresDB(),
			config.PostgresPort(),
		)
		gormLogger := logger.New(
			log.New(os.Stdout, "", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Warn,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  false,
			},
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
		if err != nil {
			return err
		}

		sqlDB, err := db.DB()
		if err != nil {
			return fmt.Errorf("failed to get underlying sql.DB: %v", err)
		}
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(5)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
		sqlDB.SetConnMaxIdleTime(3 * time.Minute)

		GormDB = db
	}
	return nil
}

func GetGormDB() (*gorm.DB, error) {
	err := ConnectToGorm()
	if err != nil {
		return nil, err
	}
	return GormDB, nil
}
