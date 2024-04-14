package models

// CREATE TABLE links (
// 	id SERIAL PRIMARY KEY,
// 	long_url TEXT NOT NULL,
// 	short_code TEXT NOT NULL,
// 	clicks INT DEFAULT 0,
// 	active BOOLEAN DEFAULT TRUE,
// 	user_id UUID NOT NULL,
// 	-- store time in epoch format
// 	created_at BIGINT NOT NULL,
// 	updated_at BIGINT NOT NULL,
// 	expires_at BIGINT NOT NULL
// );

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
