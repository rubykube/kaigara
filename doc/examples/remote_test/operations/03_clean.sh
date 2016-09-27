#!/bin/sh

sleep 3

echo "Cleaning up containers..."
docker stop etcd0
docker rm etcd0
