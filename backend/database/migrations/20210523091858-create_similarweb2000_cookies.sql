
-- +migrate Up
CREATE TABLE IF NOT EXISTS similarweb_2000_cookies (
    id INT(11) NOT NULL AUTO_INCREMENT,
    sim2000_page_id INT(11) NOT NULL,
    cookie_domain VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    KEY fk_sim2000_page_id (sim2000_page_id),
    CONSTRAINT fk_sim2000_page_id FOREIGN KEY (sim2000_page_id) REFERENCES similarweb_2000_pages (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS similarweb_2000_cookies;
