#!/bin/bash
set -e

kubectl apply -f kubernetes/reddit-auth.yaml
kubectl apply -f kubernetes/jobs.yaml
