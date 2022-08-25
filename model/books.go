package model

import "time"

type Book struct {
	ID          int        `json:"id"`
	Author      string     `json:"author"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Rating      *int       `json:"rating,omitempty"`
	CreatedAt   *time.Time `json:"created_at"`
	ModifiedAt  *time.Time `json:"modified_at"`
	Disabled    bool       `json:"disabled"`
	DisabledAt  *time.Time `json:"disabled_at,omitempty"`
}

type Books struct {
	Books []*Book `json:"books"`
}
