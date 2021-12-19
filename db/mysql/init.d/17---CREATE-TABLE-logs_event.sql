-- +migrate Up
CREATE TABLE `logs_event` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `task_id` int NOT NULL,
  `condition_id` int DEFAULT NULL,
  `time_on_page` int NOT NULL DEFAULT '0',
  `serp_page` int NOT NULL DEFAULT '0',
  `serp_rank` int NOT NULL DEFAULT '0',
  `is_visible` tinyint(1) NOT NULL DEFAULT '0',
  `event` varchar(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
