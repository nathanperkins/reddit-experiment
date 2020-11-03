#!/bin/bash

docker build -f docker/top_post_set/Dockerfile -t localhost:5000/reddit-job .
docker push localhost:5000/reddit-job
