
-- +migrate Up
CREATE TABLE IF NOT EXISTS similarweb_pages (
    id INT(11) NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    url VARCHAR(512) NOT NULL,
    icon_path VARCHAR(512) NOT NULL,
    category INT(11),
    PRIMARY KEY (id),
    KEY fk_similarweb_category_id (category),
    CONSTRAINT fk_similarweb_category_id FOREIGN KEY (category) REFERENCES similarweb_categories (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- +migrate Down
DROP TABLE IF EXISTS similarweb_pages;
