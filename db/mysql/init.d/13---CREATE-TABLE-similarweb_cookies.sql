-- +migrate Up
CREATE TABLE `similarweb_cookies` (
  `id` int NOT NULL AUTO_INCREMENT,
  `page_id` int NOT NULL,
  `domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_similarweb_page_id` (`page_id`),
  CONSTRAINT `fk_sim2000_page_id` FOREIGN KEY (`page_id`) REFERENCES `similarweb_pages` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;