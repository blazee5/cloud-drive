-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS files(
    id SERIAL,
    name VARCHAR(255) UNIQUE NOT NULL,
    download_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE files;
-- +goose StatementEnd
