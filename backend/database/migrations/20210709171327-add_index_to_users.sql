
-- +migrate Up
ALTER TABLE users ADD UNIQUE INDEX idx_users_uid(uid);

-- +migrate Down
ALTER TABLE users DROP UNIQUE INDEX idx_users_uid;
