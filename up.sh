#!/bin/sh

if [ "$1" = "--staging" ]; then
    # 도커레지스트리(private? public?) 에서 풀 한뒤 deploy
    echo "blank"
elif [ "$1" = "--develop" ]; then
    docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d --build $2
else
    echo "Usage: bash up.sh [--dev| --staging]"
fi
