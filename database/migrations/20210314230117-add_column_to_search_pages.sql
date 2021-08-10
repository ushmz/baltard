
-- +migrate Up
ALTER TABLE search_pages
    ADD `item_id` VARCHAR(255) NOT NULL;

-- +migrate Down
ALTER TABLE search_pages DROP `item_id`;
