
-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks (
    id INT(11) NOT NULL AUTO_INCREMENT,
    query VARCHAR(200) NOT NULL,
    title VARCHAR(200),
    description LONGTEXT,
    author_id INT(11) NOT NULL,
    search_url VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY fk_tasks_author_id (author_id),
    CONSTRAINT fk_tasks_author_id FOREIGN KEY (author_id) REFERENCES authors (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS tasks;