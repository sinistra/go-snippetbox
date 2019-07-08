-- +goose Up
CREATE TABLE snippets
(
    id      INTEGER      NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title   VARCHAR(255) NOT NULL,
    content TEXT         NOT NULL,
    created DATETIME     NOT NULL,
    expires DATETIME     NOT NULL
);
CREATE INDEX idx_snippets_created ON snippets (created);
-- +goose Down
DROP INDEX idx_snippets_created;
DROP TABLE snippets;
