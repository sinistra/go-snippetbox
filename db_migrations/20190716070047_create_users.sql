-- +goose Up
CREATE TABLE users
(
    id       INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name     VARCHAR(255) NOT NULL,
    email    VARCHAR(255) NOT NULL UNIQUE,
    password CHAR(60)     NOT NULL,
    created  DATETIME     NOT NULL
);

-- +goose Down
DROP TABLE users