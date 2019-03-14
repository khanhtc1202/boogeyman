#!/usr/bin/env bash

apiContainer="boogeyman-rest"
viewContainer="boogeyman-view"

docker stop ${apiContainer}
docker rm ${apiContainer}

docker run \
    -d \
    -p 3000:3000 \
    --name ${apiContainer} \
    khanhtc3010/boogeyman-rest:latest

docker stop ${viewContainer}
docker rm ${viewContainer}

docker run \
    -d \
    -p 3300:80 \
    --name ${viewContainer} \
    khanhtc3010/boogeyman-view:latest
