#!/usr/bin/env bash

echo "stop containers";
docker container stop gd_db.mf gathering.mf analyser.mf
echo "drop containers"
docker rm -v gd_db.mf gathering.mf analyser.mf

FILE_HASH=$(git rev-parse HEAD)
export GIT_HASH=$FILE_HASH

echo "RUN docker-compose-dev.yml "
serviceList="gd_db gathering_app analyser_app"
#serviceList="gd_db dc_tarantool"
echo "RUNNING SERVICES: $serviceList"
docker-compose -f docker-compose.yml pull
docker-compose -f docker-compose.yml up --build ${serviceList}