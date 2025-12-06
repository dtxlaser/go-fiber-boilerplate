package database

import (
    "app/src/config"
    "app/src/utils"
    "fmt"
    "os"
    "strconv"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
    var dsn string

    // Prefer DATABASE_URL if set (Supabase)
    databaseURL := os.Getenv("DATABASE_URL")
    if databaseURL != "" {
        dsn = databaseURL + " sslmode=require"
    } else {
        // fallback to individual env vars
        port, err := strconv.Atoi(os.Getenv("DB_PORT"))
        if err != nil {
            utils.Log.Fatalf("Invalid DB_PORT: %v", err)
        }

        dsn = fmt.Sprintf(
            "host=%s user=%s password=%s dbname=%s port=%d sslmode=require TimeZone=Asia/Shanghai",
            os.Getenv("DB_HOST"),
            os.Getenv("DB_USER"),
            os.Getenv("DB_PASSWORD"),
            os.Getenv("DB_NAME"),
            port,
        )
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger:                 logger.Default.LogMode(logger.Info),
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
        TranslateError:         true,
    })
    if err != nil {
        utils.Log.Fatalf("Failed to connect to database: %+v", err)
    }

    sqlDB, errDB := db.DB()
    if errDB != nil {
        utils.Log.Fatalf("Failed to get DB instance: %+v", errDB)
    }

    // Connection pooling
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(60 * time.Minute)

    return db
}
