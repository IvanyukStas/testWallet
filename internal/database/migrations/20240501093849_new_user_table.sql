-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) unique NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS wallets (
 id SERIAL PRIMARY KEY,
 balance VARCHAR(255) NOT NULL DEFAULT 0,
 user_id int,
 FOREIGN KEY (user_id)
    REFERENCES users(id) ON DELETE CASCADE,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS wallets;
-- +goose StatementEnd
