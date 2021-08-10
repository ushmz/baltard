
-- +migrate Up
ALTER TABLE users
	ADD `upload_file_path` varchar(512) NOT NULL DEFAULT '',
	ADD `is_task_ready` tinyint(1) NOT NULL DEFAULT 0;

-- +migrate Down
ALTER TABLE users
	DROP `upload_file_path`,
	DROP `is_task_ready`;
