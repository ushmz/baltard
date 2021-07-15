
-- +migrate Up
ALTER TABLE search_pages
    ADD `good_entity` VARCHAR(255) NOT NULL,
    ADD `bad_entity` VARCHAR(255) NOT NULL;

-- +migrate Down
ALTER TABLE search_pages
    DROP `good_entity`,
    DROP `bad_entity`;