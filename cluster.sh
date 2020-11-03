#!/bin/bash

# start up a container running the local docker image registry
reg_name='kind-registry'
reg_port='5000'
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
  docker run \
    -d --restart=always -p "${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

# start up the cluster in kind
cluster_name='reddit-experiment'
clusters="$(kind get clusters)"
if [[ "$clusters" != *"$cluster_name"* ]]; then
    kind create cluster --name "$cluster_name" --config=kubernetes/kind/cluster.yaml
fi

# connect the registry container to the kind network
network="$(docker network inspect kind --format '{{range .Containers}}{{.Name}} {{end}}')"
if [[ "$network" != *"$reg_name"* ]]; then
    docker network connect kind "$reg_name"
fi
