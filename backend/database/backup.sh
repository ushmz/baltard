#!/bin/bash

if [ $# != 2 ]; then
    echo "Wrong number of arguments. (backup.sh [containerID] [db password])"
    exit 1
fi

GIT_MYSQL=./backup/`date "+%Y-%m-%d--%H-%M-%S"`---koolhaas@localhost
mkdir -p $GIT_MYSQL
for T in `docker exec ${1} mysql -u koolhaas -pkoolhaas -N -B -e 'show tables from koolhaas'`;
do
    echo "--- Backing up $T ---"
    docker exec ${1} mysqldump --skip-comments --no-tablespaces  -u koolhaas -p${2} -d -n koolhaas $T > $GIT_MYSQL/CREATE_TABLE---$T.sql
    docker exec ${1} mysqldump --skip-comments --no-tablespaces  -u koolhaas -p${2} -t koolhaas $T > $GIT_MYSQL/INSERT_INTO---$T.sql
done;

# docker exec ${1} mysqldump -u koolhaas -p${2} koolhaas -n --no-tablespaces > `date "+%Y-%m-%d--%H-%M-%S"`---koolhaas@localhost
