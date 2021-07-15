import csv
from elasticsearch import helpers, Elasticsearch


def csv_reader(file_name):
    es = Elasticsearch([{'host': 'localhost', 'port': 9200}])
    with open(file_name, 'r') as outfile:
        reader = csv.DictReader(outfile)
        helpers.bulk(
            es, reader, index="jalan-hotel-version-may", doc_type="type")


csv_reader('/Users/kirohi/jalan/hotel_first.csv')
print('bulked')
