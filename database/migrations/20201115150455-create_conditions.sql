-- +migrate Up
CREATE TABLE IF NOT EXISTS conditions (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `condition` VARCHAR(255) NOT NULL,
    `author_id` INT(11) NOT NULL,
    PRIMARY KEY (`id`),
    KEY fk_conditions_author_id (`author_id`),
    CONSTRAINT fk_conditions_author_id FOREIGN KEY (`author_id`) REFERENCES authors (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS conditions;