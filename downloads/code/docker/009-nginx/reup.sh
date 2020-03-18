#!/bin/bash
docker-compose down
docker-compose build
docker images |grep "<none>"|awk '{print $3}'|xargs docker rmi
docker-compose up -d
docker-compose scale goweb=5  