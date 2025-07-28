package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/yudha-Dlesmana/fakeAPI/docs"
	"github.com/yudha-Dlesmana/fakeAPI/routers"

	swagger "github.com/swaggo/fiber-swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		println("🔥 Request masuk ke path:", c.Path())
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,OPTIONS",
	}))

	app.Get("/swagger/*", swagger.WrapHandler)

	routers.SetupRoutes(app)

	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is required but not set")
	}

	log.Fatal(app.Listen(":" + port))
}
