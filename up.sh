#!/bin/sh

if [ $1 == "--staging"]; then
    # 도커레지스트리(private? public?) 에서 풀 한뒤 deploy
else
    docker-compose -f ./docker-compose.yml -f ./docker-compose.dev.yml up -d --build
fi