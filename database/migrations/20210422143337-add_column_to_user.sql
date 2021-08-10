
-- +migrate Up
ALTER TABLE users
	ADD `generated_secret` varchar(12) NOT NULL;

-- +migrate Down
ALTER TABLE users
	DROP `generated_secret`;
