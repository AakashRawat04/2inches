package models

import (
	"time"
)

type Link struct {
	ID        int       `json:"id" db:"id"`
	LongURL   string    `json:"long_url" db:"long_url"`
	ShortCode string    `json:"short_code" db:"short_code"`
	Clicks    int       `json:"clicks" db:"clicks"`
	Active    bool      `json:"active" db:"active"`
	UserID    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

type CreateLinkRequest struct {
	LongURL   string    `json:"long_url"`
	ShortCode string    `json:"short_code"`
	UserID    string    `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}
