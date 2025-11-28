package models

import "time"

type User struct {
	ID        string
	Nickname  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
