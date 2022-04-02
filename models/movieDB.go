package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

/* Get genres, if any */
func (m *DBModel) getGenres(ctx context.Context, id int) map[int]string {
	query := `SELECT 
				 g.genre_name
			 FROM 
			 	genres g
			 LEFT JOIN movies_genres mg ON mg.genre_id = g.id
			 WHERE 
			 	mg.movie_id = $1`
	genreRows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil
	}
	defer genreRows.Close()

	genres := make(map[int]string)
	i := 1 // Init genre id
	for genreRows.Next() {
		var mg MovieGenre
		err := genreRows.Scan(
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil
		}

		genres[i] = mg.Genre.GenreName
		i += 1
	}

	return genres
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
	movie.MovieGenre = m.getGenres(ctx, id)

	return &movie, nil
}

/* GetAllMovies return all movies and error (if any) */
func (m *DBModel) GetAllMovies() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM movies`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*Movie

	for rows.Next() {
		var movie Movie
		err := rows.Scan(
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

		movie.MovieGenre = m.getGenres(ctx, movie.ID)

		movies = append(movies, &movie)
	}

	return movies, nil
}

/* Get list of genres */
func (m *DBModel) GetGenres() ([]*Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT * FROM genres`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*Genre

	for rows.Next() {
		var genre Genre
		err := rows.Scan(
			&genre.ID,
			&genre.GenreName,
			&genre.CreatedAt,
			&genre.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		genres = append(genres, &genre)
	}

	return genres, nil
}
