#!/bin/sh
set -e

docker-compose -f docker-compose.yml down
docker-compose -f docker-compose_dev.yml down