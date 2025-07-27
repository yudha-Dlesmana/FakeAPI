package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/yudha-Dlesmana/fakeAPI/docs"
	"github.com/yudha-Dlesmana/fakeAPI/routers"

	swagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
)

// @title 			Fake API
// @version 		1.0
// @description This is a fake REST API for prototyping, frontend development, or testing purposes only
// @host 				localhost:3000
// @basePath		/api
func main() {
	app := fiber.New()

	// Swagger route
	app.Get("/swagger/*", swagger.WrapHandler)

	routers.SetupRoutes(app)

	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is required but not set")
	}

	log.Fatal(app.Listen(":" + port))
}
