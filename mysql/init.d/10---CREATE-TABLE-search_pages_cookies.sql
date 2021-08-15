-- +migrate Up
CREATE TABLE `search_pages_cookies` (
  `id` int NOT NULL AUTO_INCREMENT,
  `page_id` int NOT NULL,
  `task_id` int NOT NULL,
  `cookie_domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_search_pages_cookies_page_id` (`page_id`),
  KEY `fk_search_pages_cookies_task_id` (`task_id`),
  KEY `idx_cookie_domain` (`cookie_domain`) USING BTREE,
  CONSTRAINT `fk_search_pages_cookies_page_id` FOREIGN KEY (`page_id`) REFERENCES `search_pages` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_search_pages_cookies_task_id` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;