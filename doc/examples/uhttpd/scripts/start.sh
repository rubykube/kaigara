#!/bin/bash

VOLDIR=$(mktemp -d)

mkdir $VOLDIR/var
mkdir $VOLDIR/etc

docker run --name uhttpd -d -p '9090:9090' \
  -v $VOLDIR/var:/var/lib/uhttpd \
  -v $VOLDIR/etc:/etc/uhttpd \
  uhttpd:0.0.1

echo "After you can delete folder: ${VOLDIR}"
