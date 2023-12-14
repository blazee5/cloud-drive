-- +goose Up
-- +goose StatementBegin
CREATE TABLE files(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    content_type VARCHAR(255) NOT NULL,
    download_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE files;
-- +goose StatementEnd
