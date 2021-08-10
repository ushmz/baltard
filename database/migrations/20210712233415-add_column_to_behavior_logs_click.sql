
-- +migrate Up
ALTER TABLE behavior_logs_click
    ADD is_visible TINYINT(1) NOT NULL DEFAULT 0;

-- +migrate Down
ALTER TABLE behavior_logs_click
    DROP is_visible;
