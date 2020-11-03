#!/bin/bash
set -e

kubectl apply -f kubernetes/secrets/reddit-auth.yaml
kubectl apply -f kubernetes/cronjobs/top-post.yaml
kubectl apply -f kubernetes/deployments/redis.yaml
kubectl apply -f kubernetes/services/redis.yaml
