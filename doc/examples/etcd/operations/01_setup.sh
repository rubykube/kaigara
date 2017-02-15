#!/bin/sh

IP="127.0.0.1"

docker run -d -p 8001:8001 -p 5001:5001 etcd:latest -peer-addr $IP:8001 -addr $IP:5001
