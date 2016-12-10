#!/bin/bash

set -xe

VOLDIR=$(mktemp -d)

mkdir $VOLDIR/var
mkdir $VOLDIR/etc

docker run -p '9090:9090' -d \
  -v metadata.yml:/etc/kaigara/metadata.yml \
  -v $VOLDIR/var:/var/lib/uhttpd \
  -v $VOLDIR/etc:/etc/uhttpd \
  uhttpd:0.0.1

echo "After you can delete folder: ${VOLDIR}"
