-- +goose Up
-- +goose StatementBegin
ALTER TABLE links
DROP COLUMN IF EXISTS description;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
