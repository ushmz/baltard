
-- +migrate Up
CREATE TABLE IF NOT EXISTS similarweb_cookies (
    id INT(11) NOT NULL AUTO_INCREMENT,
    page_id INT(11) NOT NULL,
    domain VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    KEY fk_similarweb_page_id (page_id),
    CONSTRAINT fk_sim2000_page_id FOREIGN KEY (page_id) REFERENCES similarweb_pages (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- +migrate Down
DROP TABLE IF EXISTS similarweb_cookies;
