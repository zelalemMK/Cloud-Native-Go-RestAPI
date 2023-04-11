-- +goose Up
-- This SQL will run when migration is applied
CREATE TABLE IF NOT EXISTS books (
    id UUID PRIMARY KEY, 
    title TEXT NOT NULL,
    author TEXT NOT NULL, 
    published_date DATE NOT NULL, 
    image_url TEXT NULL, 
    description TEXT NULL, 
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL, 
    deleted_at TIMESTAMP NULL
); 
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Rollback SQL
SELECT 'down SQL query';
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
