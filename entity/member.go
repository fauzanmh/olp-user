package entity

import (
	"database/sql"
)

// model
type Member struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Address   string        `json:"address"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
	DeletedAt sql.NullInt64 `json:"deleted_at"`
}

// --- params and rows --- //

type RegisterParams struct {
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Address   string        `json:"address"`
	CreatedAt int64         `json:"created_at"`
	UpdatedAt sql.NullInt64 `json:"updated_at"`
}
