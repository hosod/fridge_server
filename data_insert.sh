#!/bin/sh
set -e

curl -X POST -H "Content-Type:application/json" -d '{"expiration_date":"2020/07/28", "quantity":300, "fridge_id":2,"food_type_id":5}' localhost:8000/contents
curl -X POST -H "Content-Type:application/json" -d '{"expiration_date":"2020/07/28", "quantity":2, "fridge_id":1,"food_type_id":3}' localhost:8000/contents
curl -X POST -H "Content-Type:application/json" -d '{"expiration_date":"2020/07/28", "quantity":250, "fridge_id":1,"food_type_id":4}' localhost:8000/contents