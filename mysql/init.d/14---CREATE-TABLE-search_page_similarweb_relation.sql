-- +migrate Up
CREATE TABLE `search_page_similarweb_relation` (
  `page_id` int NOT NULL DEFAULT '0',
  `task_id` int NOT NULL,
  `similarweb_id` int NOT NULL DEFAULT '0',
  `idf` double DEFAULT NULL,
  KEY `idx_page_id` (`page_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;