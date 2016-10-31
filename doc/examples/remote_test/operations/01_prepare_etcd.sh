#!/bin/sh

export IP=127.0.0.1

set -xe

docker run -d \
  -p 2379:2379 \
  -p 4001:4001 \
  --name etcd0 \
  -v /usr/share/ca-certificates/:/etc/ssl/certs \
  quay.io/coreos/etcd:v2.3.0-alpha.1 \
  --listen-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001 \
  --advertise-client-urls http://0.0.0.0:2379,http://0.0.0.0:4001

bin/crypt set -plaintext /config/kaigara.json data.json
