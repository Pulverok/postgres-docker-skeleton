#!/bin/sh

# postgres recreating
DB_SCHEMA_PATH=$(pwd)/databases/studydb

set -e

echo 'RESTART CONTAINER'
docker restart study-postgresql

echo 'WAITING FOR POSTGRESQL'
sleep 5

echo 'GENERATE RECREATE SCRIPT FOR DATABASE studydb'
docker run --rm -v "${DB_SCHEMA_PATH}":/tmp -e DATA_PATH=/tmp skeleton_schema_generator

echo 'INIT NEW DATABASE'
cat "${DB_SCHEMA_PATH}"/init.sql | \
  docker-compose exec -T study-postgresql psql -h localhost -p 5432 -U postgres

echo 'CREATE SCHEMA'
cat ${DB_SCHEMA_PATH}/migrate.sql | \
  docker-compose exec -T study-postgresql psql -h localhost -p 5432 -U postgres study_db

if [ $? -eq 0 ]; then
    echo
    echo 'study-postgresql success!'
else
    echo
    echo 'study-postgresql failed!'
fi
