-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS answer
(
    id        BIGSERIAL NOT NULL PRIMARY KEY,
    quiz_id   BIGINT    NOT NULL,
    author_id BIGINT    NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS answer;

-- +goose StatementEnd
