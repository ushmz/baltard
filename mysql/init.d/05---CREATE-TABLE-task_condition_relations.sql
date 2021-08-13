-- +migrate Up
CREATE TABLE `task_condition_relations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `task_id` int NOT NULL,
  `condition_id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `group_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_relations_task_id` (`task_id`),
  KEY `fk_relations_condition_id` (`condition_id`),
  CONSTRAINT `fk_relations_condition_id` FOREIGN KEY (`condition_id`) REFERENCES `conditions` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_relations_task_id` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;