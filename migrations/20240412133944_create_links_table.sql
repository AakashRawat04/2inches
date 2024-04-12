-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    long_url TEXT NOT NULL,
    short_code TEXT NOT NULL,
    clicks INT DEFAULT 0,
    active BOOLEAN DEFAULT TRUE,
    user_id UUID NOT NULL,
    -- store time in epoch format
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    expires_at BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd
