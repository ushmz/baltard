
-- +migrate Up
CREATE TABLE similarweb_categories (
    id INT(11) PRIMARY KEY AUTO_INCREMENT,
    category VARCHAR(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS similarweb_categories;
