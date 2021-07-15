import csv
from elasticsearch import Elasticsearch as es
from elasticsearch.helpers import bulk
from elasticsearch_dsl import Document, Text, Float, Integer, connections


"""
10/11
このbulk.pyによって，elasticSearchにインデックス用データを投げる
elasticsearch_dslによって，es_index.csvのカラムを型付けする
"""

connections.create_connection(hosts=["localhost:9222"])


class Doc(Document):
    Id = Text()
    TotalScore = Float()
    BathScore = Float()
    BreakfastScore = Float()
    DinnerScore = Float()
    ServiceScore = Float()
    RoomScore = Float()
    PrefectureId = Integer()

    class Index:
        name = "okayama"
        doc_type = "type"


Doc.init()


def gendata(reader):
    for i, row in enumerate(reader):
        try:
            yield Doc(**row).to_dict(True)
        except:
            print(i)


def csv_reader(file_name):
    with open(file_name, 'r') as outfile:
        r = csv.DictReader(outfile)
        bulk(connections.get_connection(),
             (Doc(**row).to_dict(True) for row in r))


csv_reader('data/es_index_new.csv')
print('bulked')
