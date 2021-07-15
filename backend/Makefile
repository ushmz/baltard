build:
	docker-compose build

up:
	docker-compose up -d

deploy:
	gcloud app deploy

connect_proxy:
	./cloud_sql_proxy -instances=${CONNECT_NAME} -credential_file=key.json 