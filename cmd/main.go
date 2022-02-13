package main

import (
	"fmt"
	"log"

	"github.com/arrrden/cass-docker/pkg/db"
	"github.com/arrrden/cass-docker/pkg/db/migrations"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)

var port string = "5000"

func migrate(session *gocql.Session) error {
	fmt.Println("Checking for outstanding db migrations")

	m := migrations.Migrate{
		Session: session,
	}

	err := m.Run([]migrations.MigrationFunc{
		m.CreateUserAndPostTables(),
	})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	session, err := db.InitConnection()
	if err != nil {
		log.Fatalf("Failed to create db connection: %s", err)
	}

	err = migrate(session)
	if err != nil {
		log.Fatalf("Failed to apply migration: %s", err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	if err := app.Listen(string(":" + port)); err != nil {
		log.Fatalf("Failed to create start server")
	}
}
