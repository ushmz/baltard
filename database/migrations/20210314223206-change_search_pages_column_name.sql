
-- +migrate Up
ALTER TABLE search_pages CHANGE COLUMN `scores` `score` FLOAT;

-- +migrate Down
