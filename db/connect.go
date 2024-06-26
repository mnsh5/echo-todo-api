package db

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mnsh5/todo-api/config"
	"github.com/mnsh5/todo-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatal("Failed to parse database port:", err)
	}
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	dbName := config.Config("DB_NAME")
	log.Printf("Connection open to database %v", dbName)

	log.Println("Connection Opened to Database")
	DB.AutoMigrate(&models.Task{})
	log.Println("Database migrated successfully")
}
