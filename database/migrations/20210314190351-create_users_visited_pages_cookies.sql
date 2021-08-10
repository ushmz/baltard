
-- +migrate Up
CREATE TABLE IF NOT EXISTS users_visited_pages_cookies (
    id INT(11) NOT NULL AUTO_INCREMENT,
    visited_page_id INT(11) NOT NULL,
    cookie_domain VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    KEY fk_user_visited_page_id (visited_page_id),
    CONSTRAINT fk_user_visited_page_id FOREIGN KEY (visited_page_id) REFERENCES user_visited_pages (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS users_visited_pages_cookies;
