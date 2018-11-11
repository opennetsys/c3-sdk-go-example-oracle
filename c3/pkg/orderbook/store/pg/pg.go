package pg

import (
	"database/sql"
	"log"
)

func New(opts *Options) (*Service, error) {
	db, err := sql.Open("postgres", opt.PostgresURL)
	if err != nil {
		log.Printf("err opening postgres; err: %v", err)
		return nil, err
	}

	return &Service{
		opts: opts,
		db:   db,
	}, db.Ping()
}
