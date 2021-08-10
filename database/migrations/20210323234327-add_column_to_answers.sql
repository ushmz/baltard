
-- +migrate Up
ALTER TABLE answers
    ADD evidence_url VARCHAR(255) NOT NULL;

-- +migrate Down
ALTER TABLE answers
    DROP evidence_url;
