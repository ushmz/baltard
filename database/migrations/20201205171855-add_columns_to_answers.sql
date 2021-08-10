
-- +migrate Up
ALTER TABLE answers
    ADD `uid` VARCHAR(255) NOT NULL,
    ADD `author_id` INT(11) NOT NULL,
    ADD `hotel_id` INT(11) DEFAULT NULL, 
    ADD `reason` TEXT DEFAULT NULL;

-- +migrate Down
ALTER TABLE answers
    DROP `uid`,
    DROP `author_id`,
    DROP `hotel_id`,
    DROP `reason`;
