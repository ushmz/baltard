
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_visited_pages (
    id INT(11) NOT NULL AUTO_INCREMENT,
    user_id INT(11) NOT NULL,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(512) NOT NULL,
    screenshot_path VARCHAR(512) NOT NULL,
    PRIMARY KEY (id),
    KEY fk_user_visited_pages_user_id (user_id),
    CONSTRAINT fk_user_visited_pages_user_id FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS user_visited_pages;
