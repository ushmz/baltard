build:
	docker-compose build

run:
	docker-compose up 

bulk_insert:
	python bulk.py

setup:
	elasticsearch_loader --es-host localhost:9200 --index jalan-hotel-version-may --type type csv hotel_copy_truncate_score.csv 