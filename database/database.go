package database

import (
    "fmt"
    "log"
    "os"

    "github.com/marcellinoknl/-BE-Mini-Projects---Delosaqua/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

// Dbinstance represents a database instance with a GORM DB connection.
type Dbinstance struct {
    Db *gorm.DB
}

// DB is a global variable to hold the database instance.
var DB Dbinstance

// ConnectDb establishes a connection to the PostgreSQL database and initializes the database instance.
func ConnectDb() {
    // Construct the database connection string using environment variables.
    dsn := fmt.Sprintf(
        "host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    // Attempt to open a connection to the database.
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // Configure GORM to log SQL statements in Info mode.
    })

    if err != nil {
        log.Fatal("Warning: Failed to connect to the database. Try again!. \n", err)
        os.Exit(2)
    }

    log.Println("Connected to the database")

    // Set the GORM logger to log SQL statements in Info mode.
    db.Logger = logger.Default.LogMode(logger.Info)

    log.Println("Running database migrations")

    // AutoMigrate creates or updates the database tables based on the provided models.
    db.AutoMigrate(&models.Delosfarm{}, &models.Pond{})

    // Initialize the global database instance 'DB'.
    DB = Dbinstance{
        Db: db,
    }
}
