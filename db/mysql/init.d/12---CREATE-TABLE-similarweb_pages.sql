-- +migrate Up
CREATE TABLE `similarweb_pages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `icon_path` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `category` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_similarweb_category_id` (`category`),
  CONSTRAINT `fk_similarweb_category_id` FOREIGN KEY (`category`) REFERENCES `similarweb_categories` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;