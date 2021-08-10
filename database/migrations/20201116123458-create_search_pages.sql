
-- +migrate Up
CREATE TABLE IF NOT EXISTS search_pages (
    id INT(11) NOT NULL AUTO_INCREMENT,
    title VARCHAR(512) NOT NULL,
    url VARCHAR(512) NOT NULL,
    snippet VARCHAR(512) NOT NULL,
    rank INT(11) NOT NULL,
    misc VARCHAR(255) DEFAULT NULL,
    scores FLOAT DEFAULT NULL,
    task_id INT(11) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY fk_search_pages_task_id (task_id),
    CONSTRAINT fk_search_pages_task_id FOREIGN KEY (task_id) REFERENCES tasks (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS search_pages;
