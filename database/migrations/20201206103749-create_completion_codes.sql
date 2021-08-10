
-- +migrate Up
CREATE TABLE completion_codes (
    id INT(11) NOT NULL AUTO_INCREMENT,
    uid VARCHAR(255) NOT NULL DEFAULT '',
    completion_code INT(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS completion_codes;
