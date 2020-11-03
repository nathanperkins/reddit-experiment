#!/bin/bash

docker build -f docker/top_post_set/Dockerfile -t localhost:5000/top-post-set .
docker push localhost:5000/top-post-set

docker build -f docker/top_post_get/Dockerfile -t localhost:5000/top-post-get .
docker push localhost:5000/top-post-get
