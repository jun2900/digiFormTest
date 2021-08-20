package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jun2900/digiFormTest/controllers"
	"github.com/jun2900/digiFormTest/database"
	"github.com/jun2900/digiFormTest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initMainDatabase() {
	var err error
	postgresUser := os.Getenv("DB_USER")
	postgresPassword := os.Getenv("DB_PASSWORD")
	postgresPort := os.Getenv("DB_PORT")
	postgresDatabase := os.Getenv("DB_NAME")
	postgresHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", postgresHost, postgresPort, postgresUser, postgresDatabase, postgresPassword)
	database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connection open")
	database.DBConn.AutoMigrate(&models.Lokasi{}, &models.AirwayBill{})
}

func main() {
	app := fiber.New()

	//Load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initMainDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/upload", controllers.AirwayBill)

	app.Listen(":3000")

}
