-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email TEXT CONSTRAINT unique_email UNIQUE
);
-- +goose Down
DROP TABLE users;
