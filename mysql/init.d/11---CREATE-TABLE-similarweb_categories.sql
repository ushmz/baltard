-- +migrate Up
CREATE TABLE `similarweb_categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `category` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;