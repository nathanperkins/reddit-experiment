#!/bin/bash

docker build -f docker/reddit_job/Dockerfile -t localhost:5000/reddit-job .
docker push localhost:5000/reddit-job
