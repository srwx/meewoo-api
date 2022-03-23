package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

/* GetOneMovie return one movie and error (if any) */
func (m *DBModel) GetOneMovie(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM movies WHERE id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie
	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.Runtime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// get genres, if any
	query = `SELECT 
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			 FROM 
			 	movies_genres mg
			 LEFT JOIN genres g ON mg.genre_id = g.id
			 WHERE 
			 	mg.movie_id = $1`
	genreRows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer genreRows.Close()

	var genres []MovieGenre
	for genreRows.Next() {
		var mg MovieGenre
		err := genreRows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)

		if err != nil {
			return nil, err
		}
		genres = append(genres, mg)
	}

	movie.MovieGenre = genres

	return &movie, nil
}

/* GetAllMovies return all movies and error (if any) */
func (m *DBModel) GetAllMovies() ([]*Movie, error) {
	return nil, nil
}
