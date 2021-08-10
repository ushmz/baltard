
-- +migrate Up
CREATE TABLE IF NOT EXISTS similarweb_2000_pages (
    id INT(11) NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(512) NOT NULL,
    icon_path VARCHAR(512) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS similarweb_2000_pages;
