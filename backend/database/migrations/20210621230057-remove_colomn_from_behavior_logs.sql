
-- +migrate Up
ALTER TABLE behavior_logs 
    DROP `author_id`;

-- +migrate Down
ALTER TABLE behavior_logs 
    ADD `author_id` INT(11) NOT NULL;
