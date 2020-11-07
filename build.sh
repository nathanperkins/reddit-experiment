#!/bin/bash

docker build -f docker/top_post_set/Dockerfile -t nathanperkins/top-post-set .
docker push nathanperkins/top-post-set

docker build -f docker/top_post_get/Dockerfile -t nathanperkins/top-post-get .
docker push nathanperkins/top-post-get
