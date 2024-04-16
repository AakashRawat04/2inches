package models

type Link struct {
	ID        int    `json:"id" db:"id"`
	LongURL   string `json:"long_url" db:"long_url"`
	ShortCode string `json:"short_code" db:"short_code"`
	Clicks    int    `json:"clicks" db:"clicks"`
	Active    bool   `json:"active" db:"active"`
	UserID    string `json:"user_id" db:"user_id"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
	UpdatedAt int64  `json:"updated_at" db:"updated_at"`
	ExpiresAt int64  `json:"expires_at" db:"expires_at"`
}

type CreateLinkRequest struct {
	LongURL   string `json:"long_url"`
	ShortCode string `json:"short_code"`
	UserID    string `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"`
}
