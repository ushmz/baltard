
-- +migrate Up
CREATE TABLE IF NOT EXISTS group_counts (
    `group_id` INT(11) NOT NULL,
    `author_id` INT(11) NOT NULL,
    `count` INT(11) NOT NULL DEFAULT 0,
    KEY fk_groups_author_id (`author_id`),
    CONSTRAINT fk_groups_author_id FOREIGN KEY (`author_id`) REFERENCES authors (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
-- +migrate Down
DROP TABLE IF EXISTS group_counts;