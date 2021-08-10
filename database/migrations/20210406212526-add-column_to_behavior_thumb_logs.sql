
-- +migrate Up
ALTER TABLE behavior_thumb_logs
    ADD `linked_history` INT(11) NOT NULL DEFAULT 0;

-- +migrate Down
ALTER TABLE behavior_thumb_logs DROP `linked_history`;
