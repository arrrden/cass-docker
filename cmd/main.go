package main

import (
	"fmt"
	"log"

	"github.com/arrrden/cassandra-proto/pkg/db"
	"github.com/gofiber/fiber/v2"
)

var port string = "5000"

func main() {
	if err := db.InitConnection(); err != nil {
		log.Fatalf("Failed to create db connection")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(string(":" + port))
	fmt.Print("API is ready on port: %", port)
}
