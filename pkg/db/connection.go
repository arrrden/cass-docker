package db

import (
	"github.com/gocql/gocql"
)

type DBConnection struct {
	cluster *gocql.ClusterConfig
	session *gocql.Session
}

var (
	connection DBConnection
	host       string = "127.0.0.1"
	keyspace   string = "development"
)

func InitConnection() (*gocql.Session, error) {
	connection.cluster = gocql.NewCluster(host)
	connection.cluster.Consistency = gocql.Quorum
	connection.cluster.Keyspace = keyspace

	var err error

	connection.session, err = connection.cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return connection.session, nil
}
