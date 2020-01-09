package output

import "database/sql"

// PGOutputer ...
type PGOutputer struct {
	dsn string
	db  *sql.DB
}

// Output ...
func (o PGOutputer) Output(msg string) error {
	return nil
}
