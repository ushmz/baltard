export GO111MODULE := off

DB_PORT:=3366
MYSQL_USER:=koolhaas
MYSQL_DATABASE:=koolhaas
MYSQL_PASSWORD:=koolhaas
MYSQL_CONTAINER_NAME:=koolhaas-mysql

# docker-composeで立てたMySQLコンテナのMySQLに，csvデータをimportする
# FIXME: コンテナ内でコマンド実行する必要ありなので面倒(開発環境でしか使わないため放置)
import_hotels:
	docker cp csv_seeder/ $(MYSQL_CONTAINER_NAME):tmp/
	docker exec -it $(MYSQL_CONTAINER_NAME) mysql -u $(MYSQL_USER) -p $(MYSQL_PASSWORD) 
	# mysqlに入った後以下を実行
	  LOAD DATA LOCAL INFILE 'tmp/csv_seeder/hotels.kagawa.csv'  INTO TABLE hotels FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' ESCAPED BY '' IGNORE 1 LINES \
	  (@1, @2, @3, @4, @5, @6, @7, @8, @9, @10, @11, @12, @13, @14, @15, @16, @17, @18, @19, @20, @21, @22, @23, @24, @25) \
	  SET id=@1, name=@2, url=@3, prefecture=@4, large_area=@5, small_area=@6, address=@7, hotel_type=@8, caption=@9, catch=@10, lowest_price=@11, latitude=@12, longitude=@13, created_at=@14, updated_at=@15, total_score=@16, room_score=@17, bath_score=@18, breakfast_score=@19, dinner_score=@20, service_score=@21, clean_score=@22, prefecture_id=@23, good_entity=@24, bad_entity=@25;
	#  LOAD DATA LOCAL INFILE 'tmp/csv_seeder/hotels.okayama.csv'  INTO TABLE hotels FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"' ESCAPED BY '';
