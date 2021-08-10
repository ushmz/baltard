
-- +migrate Up
ALTER TABLE users
    DROP `external_id`,
	DROP `upload_file_path`,
	DROP `is_task_ready`;

-- +migrate Down
ALTER TABLE users
    ADD `external_id` VARCHAR(255) NOT NULL;
	ADD `upload_file_path` varchar(512) NOT NULL DEFAULT '',
	ADD `is_task_ready` tinyint(1) NOT NULL DEFAULT 0;
