package storage

import "database/sql"

type Book struct {
	ID          int
	Author      string
	Title       string
	Description string
	Rating      *int
	CreatedAt   sql.NullTime
	ModifiedAt  sql.NullTime
	Disabled    bool
	DisabledAt  sql.NullTime
}
