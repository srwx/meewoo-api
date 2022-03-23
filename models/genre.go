package models

import "time"

type Genre struct {
	ID        int       `json:"-"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
