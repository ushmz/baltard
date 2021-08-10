
-- +migrate Up
ALTER TABLE tasks ADD `prefecture_id` INT(11);
-- +migrate Down
ALTER TABLE tasks DROP `prefecture_id`;