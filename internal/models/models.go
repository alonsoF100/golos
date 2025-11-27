package models

import "time"

// TODO описать модели данных, с которыми сможем работать во всех слоях

type User struct {
	ID string `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
