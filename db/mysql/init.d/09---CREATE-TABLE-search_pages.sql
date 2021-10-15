-- +migrate Up
CREATE TABLE `search_pages` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `snippet` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `task_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_search_pages_task_id` (`task_id`),
  CONSTRAINT `fk_search_pages_task_id` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;