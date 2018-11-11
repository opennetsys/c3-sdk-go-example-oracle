package pg

import "database/sql"

type Options struct {
	PostgresURL string
}

type Service struct {
	opts *Options
	db   *sql.DB
}
