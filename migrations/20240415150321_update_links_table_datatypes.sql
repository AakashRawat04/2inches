-- +goose Up
-- +goose StatementBegin

-- Drop the id column
ALTER TABLE links DROP COLUMN id;

-- Make short_code column the primary key
ALTER TABLE links ADD CONSTRAINT links_short_code_pk PRIMARY KEY (short_code);

-- Create temporary columns with timestamp type
ALTER TABLE links ADD COLUMN created_at_temp TIMESTAMP WITH TIME ZONE;
ALTER TABLE links ADD COLUMN updated_at_temp TIMESTAMP WITH TIME ZONE;
ALTER TABLE links ADD COLUMN expires_at_temp TIMESTAMP WITH TIME ZONE;

-- Update temporary columns with data from existing columns
UPDATE links SET created_at_temp = TO_TIMESTAMP(created_at)::TIMESTAMP WITH TIME ZONE;
UPDATE links SET updated_at_temp = TO_TIMESTAMP(updated_at)::TIMESTAMP WITH TIME ZONE;
UPDATE links SET expires_at_temp = TO_TIMESTAMP(expires_at)::TIMESTAMP WITH TIME ZONE;

-- Drop existing columns
ALTER TABLE links DROP COLUMN created_at;
ALTER TABLE links DROP COLUMN updated_at;
ALTER TABLE links DROP COLUMN expires_at;

-- Rename temporary columns
ALTER TABLE links RENAME COLUMN created_at_temp TO created_at;
ALTER TABLE links RENAME COLUMN updated_at_temp TO updated_at;
ALTER TABLE links RENAME COLUMN expires_at_temp TO expires_at;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Reverse the changes in the Down migration
ALTER TABLE links DROP CONSTRAINT links_short_code_pk;
ALTER TABLE links ADD COLUMN id SERIAL PRIMARY KEY;

-- Create temporary columns with bigint type
ALTER TABLE links ADD COLUMN created_at_temp BIGINT;
ALTER TABLE links ADD COLUMN updated_at_temp BIGINT;
ALTER TABLE links ADD COLUMN expires_at_temp BIGINT;

-- Update temporary columns with data from existing columns
UPDATE links SET created_at_temp = EXTRACT(EPOCH FROM created_at)::BIGINT;
UPDATE links SET updated_at_temp = EXTRACT(EPOCH FROM updated_at)::BIGINT;
UPDATE links SET expires_at_temp = EXTRACT(EPOCH FROM expires_at)::BIGINT;

-- Drop existing columns
ALTER TABLE links DROP COLUMN created_at;
ALTER TABLE links DROP COLUMN updated_at;
ALTER TABLE links DROP COLUMN expires_at;

-- Rename temporary columns
ALTER TABLE links RENAME COLUMN created_at_temp TO created_at;
ALTER TABLE links RENAME COLUMN updated_at_temp TO updated_at;
ALTER TABLE links RENAME COLUMN expires_at_temp TO expires_at;

-- +goose StatementEnd
