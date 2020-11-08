#!/bin/bash
set -e

kubectl apply -f kubernetes/secret/reddit-auth.yaml
kubectl apply -f kubernetes/deployment/redis.yaml
kubectl apply -f kubernetes/service/redis.yaml
kubectl apply -f kubernetes/cronjob/top-post-set.yaml
kubectl apply -f kubernetes/deployment/top-post-get.yaml
kubectl apply -f kubernetes/service/top-post-get.yaml
kubectl apply -f kubernetes/managed-certificate/primary.yaml
kubectl apply -f kubernetes/ingress/primary.yaml
