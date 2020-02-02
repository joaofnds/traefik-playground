#!/usr/bin/env sh

docker-compose up -d --scale nginx=10 --scale whoami=10 --scale dumper=10
