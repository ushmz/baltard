
-- +migrate Up
ALTER TABLE users
    ADD `external_id` VARCHAR(255) NOT NULL;

-- +migrate Down
ALTER TABLE users DROP `external_id`;