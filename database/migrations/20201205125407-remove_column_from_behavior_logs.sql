
-- +migrate Up
ALTER TABLE behavior_logs 
    DROP `refocus_count`,
    DROP `click_count`;

-- +migrate Down
ALTER TABLE behavior_logs 
    ADD `refocus_count`,
    ADD `click_count`;