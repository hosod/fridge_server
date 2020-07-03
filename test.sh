#!/bin/sh
set -e

curl localhost:8000/contents/fridge?fid=1
echo '\n'
curl localhost:8000/contents/fridge?fid=2
echo '\n'
curl localhost: