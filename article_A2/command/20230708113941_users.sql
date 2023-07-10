-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id      INTEGER NOT NULL,
    name    VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
