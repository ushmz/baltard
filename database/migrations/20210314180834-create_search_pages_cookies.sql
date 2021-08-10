
-- +migrate Up
CREATE TABLE IF NOT EXISTS search_pages_cookies (
    id INT(11) NOT NULL AUTO_INCREMENT,
    page_id INT(11) NOT NULL,
    task_id INT(11) NOT NULL,
    cookie_domain VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    KEY fk_search_pages_cookies_page_id (page_id),
    CONSTRAINT fk_search_pages_cookies_page_id FOREIGN KEY (page_id) REFERENCES search_pages (id),
    KEY fk_search_pages_cookies_task_id (task_id),
    CONSTRAINT fk_search_pages_cookies_task_id FOREIGN KEY (task_id) REFERENCES tasks (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS search_pages_cookies;
