package migrations

import (
	"github.com/gocql/gocql"
)

type Migrate struct {
	Session *gocql.Session
}

type MigrationFunc func() error

func (m *Migrate) Run(migrations []MigrationFunc) error {
	for _, x := range migrations {
		if err := x(); err != nil {
			return err
		}
	}

	return nil
}
