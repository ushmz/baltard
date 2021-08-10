
-- +migrate Up
ALTER TABLE task_condition_relations 
    ADD `author_id` INT(11), 
    ADD `group_id` INT(11);

-- +migrate Down
ALTER TABLE task_condition_relations 
    DROP `author_id`,
    DROP `group_id`;