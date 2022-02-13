package handlerscommon

import (
	"time"

	"github.com/gocql/gocql"
)

type Data struct {
	Id        string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Keyspace  string
}

func UUID() (string, error) {
	uuid, err := gocql.RandomUUID()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}
