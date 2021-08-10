
-- +migrate Up
ALTER TABLE behavior_logs
    ADD `task_id` INT(11) NOT NULL,
    ADD `condition_id` INT(11) DEFAULT NULL,
    ADD `group_id` INT(11) NOT NULL;

-- +migrate Down
ALTER TABLE behavior_logs
    DROP `task_id`,
    DROP `condition_id`,
    DROP `group_id`;
