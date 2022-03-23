package models

import "database/sql"

/* Models is the wrapper for database */
type Models struct {
	DB DBModel
}

/* NewModels return Models with db pool */
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}
