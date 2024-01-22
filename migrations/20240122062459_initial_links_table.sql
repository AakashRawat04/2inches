-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
    long_url VARCHAR(255) NOT NULL,
    short_code VARCHAR(10) PRIMARY KEY UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    last_updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    author VARCHAR(50),
    views INT DEFAULT 0,
    is_active BOOLEAN DEFAULT true,
    expiration_date TIMESTAMPTZ,
    description TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd


