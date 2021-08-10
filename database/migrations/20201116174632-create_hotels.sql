-- +migrate Up
CREATE TABLE IF NOT EXISTS hotels (
    id INT(11) NOT NULL,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    prefecture VARCHAR(255) NOT NULL,
    prefecture_id INT(11) NOT NULL,
    large_area VARCHAR(255) NOT NULL,
    small_area VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    hotel_type VARCHAR(255) NOT NULL,
    caption TEXT,
    catch TEXT,
    lowest_price INT(11) DEFAULT NULL,
    latitude DOUBLE DEFAULT NULL,
    longitude DOUBLE DEFAULT NULL,
    total_score DOUBLE  NOT NULL,
    room_score DOUBLE  NOT NULL,
    bath_score DOUBLE  NOT NULL,
    breakfast_score DOUBLE  NOT NULL,
    dinner_score DOUBLE  NOT NULL,
    service_score DOUBLE  NOT NULL,
    clean_score DOUBLE  NOT NULL,
    good_entity VARCHAR(255) DEFAULT NULL,
    bad_entity VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY hotels_prefecture_idx (prefecture),
    KEY hotels_large_area_idx (large_area),
    KEY hotels_small_area_idx (small_area),
    KEY hotels_type_idx (hotel_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS hotels;