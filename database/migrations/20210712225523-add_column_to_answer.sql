
-- +migrate Up
ALTER TABLE answers
    ADD answer VARCHAR(255) NOT NULL;

-- +migrate Down
ALTER TABLE answers
    DROP answer;
