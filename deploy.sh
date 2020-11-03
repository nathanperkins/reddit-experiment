#!/bin/bash
set -e

kubectl delete -f kubernetes/secrets/reddit-auth.yaml || true
kubectl delete -f kubernetes/cronjobs/top-post.yaml || true

kubectl apply -f kubernetes/secrets/reddit-auth.yaml
kubectl apply -f kubernetes/cronjobs/top-post.yaml
