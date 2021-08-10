
-- +migrate Up
CREATE TABLE IF NOT EXISTS behavior_logs_click  (
    id INT(11) NOT NULL AUTO_INCREMENT,
    uid VARCHAR(200) NOT NULL,
    task_id INT(11) NOT NULL,
    condition_id INT(11) DEFAULT NULL,
    time_on_page INT(11) NOT NULL DEFAULT 0,
    serp_page INT(11) NOT NULL DEFAULT 0,
    serp_rank INT(11) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS behavior_logs_click;
