-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    name TEXT NOT NULL,
    password VARCHAR(256) NOT NULL,
    api_key VARCHAR(64) UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;