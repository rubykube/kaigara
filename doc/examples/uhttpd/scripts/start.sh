#!/bin/bash

CURDIR="$(pwd)"
docker run --name uhttpd -d -p '9090:9090' \
  -v $CURDIR/mnt/var/lib/uhttpd:/var/lib/uhttpd \
  -v $CURDIR/mnt/etc/uhttpd:/etc/uhttpd \
  uhttpd:0.0.1
