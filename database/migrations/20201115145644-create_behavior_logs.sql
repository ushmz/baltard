
-- +migrate Up
CREATE TABLE IF NOT EXISTS behavior_logs  (
    id VARCHAR(200) NOT NULL,
    uid VARCHAR(200) NOT NULL,
    author_id VARCHAR(200) NOT NULL,
    click_count INT(11) DEFAULT 0,
    time_on_page INT(11) NOT NULL DEFAULT 0,
    referrer VARCHAR(512),
    refocus_count INT(11),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    url VARCHAR(512) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS behavior_logs;