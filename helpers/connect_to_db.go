package helpers

import (
    "inate_payment_service/config"
    "context"
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

)

var DB *sql.DB
var GormDB *gorm.DB

func ConnectToPostgres(ctx context.Context) error {
    if DB == nil {
        connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
            config.PostgresUser(),
            config.PostgresPassword(),
            config.PostgresHost(),
            config.PostgresDB(),
        )
        var err error
        DB, err = sql.Open("postgres", connStr)
        if err != nil {
            return err
        }
        // Optionally ping to verify connection
        if err = DB.PingContext(ctx); err != nil {
            return err
        }
    }
    return nil
}

func GetRawDB() (*sql.DB, error) {
    err := ConnectToPostgres(context.Background())
    if err != nil {
        return nil, err
    }
    return DB, nil
}

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
		fmt.Println("Connecting to Gorm with DSN:", dsn)
        db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            return err
        }
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
