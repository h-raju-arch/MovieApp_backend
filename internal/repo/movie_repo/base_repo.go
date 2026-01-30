package movierepo

import "database/sql"

type MovieRepo struct {
	db *sql.DB
}

func NewMovieRepo(db *sql.DB) *MovieRepo {
	return &MovieRepo{db: db}
}
