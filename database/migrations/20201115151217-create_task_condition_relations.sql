
-- +migrate Up
CREATE TABLE IF NOT EXISTS task_condition_relations (
    id INT(11) NOT NULL AUTO_INCREMENT,
    task_id INT(11) NOT NULL,
    condition_id INT(11) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY fk_relations_task_id (task_id),
    KEY fk_relations_condition_id (condition_id),
    CONSTRAINT fk_relations_task_id FOREIGN KEY (task_id) REFERENCES tasks (id),
    CONSTRAINT fk_relations_condition_id FOREIGN KEY (condition_id) REFERENCES conditions (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS task_condition_relations;