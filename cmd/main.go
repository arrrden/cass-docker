package main

import (
	"fmt"
	"log"

	"github.com/arrrden/cass-docker/pkg/db"
	"github.com/arrrden/cass-docker/pkg/db/migrations"
	"github.com/arrrden/cass-docker/pkg/handlers"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)

const (
	port     string = "5000"
	keyspace string = "development"
)

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

	defer session.Close()

	err = migrate(session)
	if err != nil {
		log.Fatalf("Failed to apply migration: %s", err)
	}
	app := fiber.New()

	userHandler := handlers.UserHandler{
		Keyspace: keyspace,
	}

	app.Get("/api/user/:id", userHandler.Get)
	app.Post("/api/user", userHandler.Post)

	if err := app.Listen(string(":" + port)); err != nil {
		log.Fatalf("Failed to create start server")
	}
}
