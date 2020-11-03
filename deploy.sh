#!/bin/bash
set -e

kubectl delete -f kubernetes/secrets/reddit-auth.yaml || true
kubectl delete -f kubernetes/cronjobs/top-post.yaml || true
kubectl delete -f kubernetes/deployments/redis.yaml || true
kubectl delete -f kubernetes/services/redis.yaml || true

kubectl apply -f kubernetes/secrets/reddit-auth.yaml
kubectl apply -f kubernetes/cronjobs/top-post.yaml
kubectl apply -f kubernetes/deployments/redis.yaml
kubectl apply -f kubernetes/services/redis.yaml
