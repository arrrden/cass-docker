package db

// God I want this to be a generic so bad...
// Query[K interface](query string, values ...interface{}) (K)
func Query(query string, values ...interface{}) error {
	err := connection.session.Query(query).Bind(values...).Exec()
	if err != nil {
		return err
	}

	return nil
}
