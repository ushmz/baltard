
-- +migrate Up
ALTER TABLE answers DROP evidence_url;

-- +migrate Down
ALTER TABLE answers
    ADD evidence_url VARCHAR(255) NOT NULL;
