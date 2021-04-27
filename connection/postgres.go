package connection

import (
	"database/sql"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Connect() error {
	// TODO
}
