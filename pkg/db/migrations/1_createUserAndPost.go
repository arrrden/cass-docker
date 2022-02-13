package migrations

import (
	"fmt"

	"github.com/gocql/gocql"
)

func createUserTable(s *gocql.Session) error {
	query := `
		CREATE TABLE IF NOT EXISTS development.user ( 
			id uuid PRIMARY KEY, 
			createdAt timestamp, 
			updatedAt timestamp,
			name text, 
			email text,
		);
	`

	if err := s.Query(query).Exec(); err != nil {
		return err
	}

	return nil
}

func createPostTable(s *gocql.Session) error {
	query := `
		CREATE TABLE IF NOT EXISTS development.post ( 
			id uuid PRIMARY KEY, 
			createdAt timestamp, 
			updatedAt timestamp,
		);
	`

	if err := s.Query(query).Exec(); err != nil {
		return err
	}

	return nil
}

func (m *Migrate) CreateUserAndPostTables() MigrationFunc {
	return func() error {
		var err error
		err = createUserTable(m.Session)
		if err != nil {
			return err
		}
		err = createPostTable(m.Session)
		if err != nil {
			return err
		}

		fmt.Println("User and Post tables ready")
		return nil
	}
}
