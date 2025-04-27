-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    title TEXT NOT NULL,
    user_id BIGINT NOT NULL,
    tags TEXT[] DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    version INT DEFAULT 1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
