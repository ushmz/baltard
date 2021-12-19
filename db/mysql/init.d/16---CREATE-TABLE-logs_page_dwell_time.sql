-- +migrate Up
CREATE TABLE `logs_page_dwell_time` (
  `user_id` int NOT NULL,
  `task_id` int NOT NULL,
  `page_id` int NOT NULL,
  `condition_id` int DEFAULT NULL,
  -- Where does this page exist on search result pages. (in 10 pages)
  -- `page` int NOT NULL,
  -- Were does this page exist on the $`page`th search result page.
  -- `rank` int NOT NULL,
  `time_on_page` int NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`task_id`, `page_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
