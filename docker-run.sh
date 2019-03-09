#!/usr/bin/env bash

container="boogeyman-rest"

docker stop ${container}
docker rm ${container}

docker run \
    -d \
    -p 3000:3000 \
    --name ${container} \
    khanhtc3010/boogeyman-rest:latest
