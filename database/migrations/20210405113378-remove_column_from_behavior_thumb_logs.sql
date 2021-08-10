
-- +migrate Up
ALTER TABLE behavior_thumb_logs 
    DROP `updated_at`;

-- +migrate Down
ALTER TABLE behavior_logs 
    ADD `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
;
