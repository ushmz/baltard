
-- +migrate Up
ALTER TABLE behavior_logs 
    DROP `referrer`,
    DROP `group_id`;

-- +migrate Down
ALTER TABLE behavior_logs 
    ADD `referrer`,
    ADD `group_id`;
