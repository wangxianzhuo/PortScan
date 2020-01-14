package output

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// PGOutputer ...
type PGOutputer struct {
	dsn string
	db  *sql.DB
}

// Output ...
func (o PGOutputer) Output(msg, address string) error {
	stmt, err := o.db.Prepare("INSERT INTO public.tcp_port_errors(address, msg)VALUES ($1, $2);")
	if err != nil {
		log.Println("Error: pgOutput error:", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(address, msg)
	if err != nil {
		log.Println("Error: pgOutput error:", err)
	}

	return nil
}
