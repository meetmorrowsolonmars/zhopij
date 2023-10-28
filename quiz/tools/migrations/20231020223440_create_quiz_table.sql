-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS quiz
(
    id        BIGSERIAL    NOT NULL PRIMARY KEY,
    author_id BIGINT       NOT NULL,
    title     VARCHAR(512) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS quiz;

-- +goose StatementEnd
