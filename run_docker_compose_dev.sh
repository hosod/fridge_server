#!/bin/sh
set -e

docker-compose -f docker-compose.yml stop
docker-compose -f docker-compose_dev.yml stop
docker-compose -f docker-compose_dev.yml up -d
