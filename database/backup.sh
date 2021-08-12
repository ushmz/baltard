#!/bin/bash

if [ $# != 2 ]; then
    echo "Wrong number of arguments. (backup.sh [containerID] [db password])"
    exit 1
fi

GIT_MYSQL=./backup/`date "+%Y-%m-%d--%H-%M-%S"`---baltard@localhost
mkdir -p $GIT_MYSQL
for T in `docker exec ${1} mysql -u baltard -p${2} -N -B -e 'show tables from baltard'`;
do
    echo "--- Backing up $T ---"
    docker exec ${1} mysqldump --skip-comments --no-tablespaces  -u baltard -p${2} -d -n baltard $T > $GIT_MYSQL/CREATE_TABLE---$T.sql
    docker exec ${1} mysqldump --skip-comments --no-tablespaces  -u baltard -p${2} -t baltard $T > $GIT_MYSQL/INSERT_INTO---$T.sql
done;

# docker exec ${1} mysqldump -u baltard -p${2} baltard -n --no-tablespaces > `date "+%Y-%m-%d--%H-%M-%S"`---baltard@localhost
