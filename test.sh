#!/bin/sh
set -e


curl localhost:8000/users?uid=1
echo '\n'
curl localhost:8000/fridges?fid=4
echo '\n'
curl localhost:8000/contents/fridge?fid=1
echo '\n'
curl localhost:8000/contents/fridge?fid=2
echo '\n'
curl -X POST -H "Content-Type: application/json" -d '{"expiration_date":"2020/08/05", "quantity":1.5, "fridge_id":1,"food_type_id":7}' localhost:8000/contents
curl -X POST -H "Content-Type: application/json" -d '{"expiration_date":"2020/07/28", "quantity":300, "fridge_id":1,"food_type_id":6}' localhost:8000/contents
curl -X POST -H "Content-Type: application/json" -d '{"expiration_date":"2020/07/28", "quantity":1.5, "fridge_id":1,"food_type_id":7}' localhost:8000/contents
curl -X POST -H "Content-Type: application/json" -d '{"expiration_date":"2020/07/28", "quantity":1.5, "fridge_id":2,"food_type_id":7}' localhost:8000/contents
curl -X POST -H "Content-Type: application/json" -d '{"expiration_date":"2020/07/28", "quantity":300, "fridge_id":2,"food_type_id":6}' localhost:8000/contents
curl -X POST -H "Content-Type: application/json" -d '{"expiration_date":"2020/07/28", "quantity":300, "fridge_id":5,"food_type_id":6}' localhost:8000/contents

echo '\n'

curl localhost:8000/contents/fridge?fid=5
echo '\n'