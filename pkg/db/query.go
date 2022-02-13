package db

import "github.com/gocql/gocql"

func FindMany(query string, values ...interface{}) *gocql.Iter {
	qry := connection.session.Query(query).Bind(values...)
	iter := qry.Iter()

	return iter
}

func FindOne(query string, values ...interface{}) *gocql.Query {
	qry := connection.session.Query(query).Bind(values...)
	return qry
}

func Mutate(query string, values ...interface{}) error {
	qry := connection.session.Query(query).Bind(values...)
	err := qry.Exec()
	if err != nil {
		return err
	}

	return nil
}
