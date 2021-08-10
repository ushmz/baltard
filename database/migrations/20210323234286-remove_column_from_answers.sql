
-- +migrate Up
ALTER TABLE answers 
    DROP task_condition_relations_id,
    DROP hotel_id;

-- +migrate Down
ALTER TABLE answers 
    ADD task_condition_relations_id INT(11),
    ADD `hotel_id` INT(11) DEFAULT NULL;

