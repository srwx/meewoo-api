package models

import "database/sql"

type DBModel struct {
	DB *sql.DB
}

/* GetOneMovie return one movie and error (if any) */
func (m *DBModel) GetOneMovie(id int) (*Movie, error) {
	return nil, nil
}

/* GetAllMovies return all movies and error (if any) */
func (m *DBModel) GetAllMovies() ([]*Movie, error) {
	return nil, nil
}
