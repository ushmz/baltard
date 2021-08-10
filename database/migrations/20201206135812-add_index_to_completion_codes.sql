
-- +migrate Up
ALTER TABLE completion_codes ADD INDEX idx_completion_codes_uid(uid);
-- +migrate Dow
ALTER TABLE completion_codes DROP INDEX idx_completion_codes_uid;
