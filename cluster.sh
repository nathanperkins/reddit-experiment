#!/bin/bash
reg_name='kind-registry'
reg_port='5000'
running="$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)"
if [ "${running}" != 'true' ]; then
  docker run \
    -d --restart=always -p "${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

cluster_name="reddit-experiment"
clusters="$(kind get clusters)"
if [[ "$clusters" != *"$cluster_name"* ]]; then
    kind create cluster --name "$cluster_name" --config=kubernetes/cluster.yaml
fi
network="$(docker network inspect kind --format '{{range .Containers}}{{.Name}} {{end}}')"
if [[ "$network" != *"kind-registry"* ]]; then
    docker network connect kind kind-registry
fi
