-- +migrate Up
CREATE TABLE IF NOT EXISTS reviews (
    id INT(11) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    hotel_id INT(11) DEFAULT NULL,
    total_score int(11) DEFAULT NULL,
    room_score int(11) DEFAULT NULL,
    bath_score int(11) DEFAULT NULL,
    breakfast_score int(11) DEFAULT NULL,
    dinner_score int(11) DEFAULT NULL,
    service_score int(11) DEFAULT NULL,
    clean_score int(11) DEFAULT NULL,
    comment text,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY reviews_created_at_idx (created_at),
    KEY reviews_total_score_idx (total_score)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS reviews;