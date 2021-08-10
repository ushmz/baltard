
-- +migrate Up
ALTER TABLE tasks
    ADD `type` VARCHAR(255) NOT NULL;

-- +migrate Down
ALTER TABLE tasks DROP `type`;
