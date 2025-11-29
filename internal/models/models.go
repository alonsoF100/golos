package models

import "time"

type User struct {
	ID        string
	Nickname  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Election struct {
	ID          string
	UserID      string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type VoteVariant struct {
	ID         string
	ElectionID string
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Vote struct {
	ID        string
	VariantID string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
