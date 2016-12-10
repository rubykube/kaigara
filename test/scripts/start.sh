#!/bin/bash

set -xe

VOLDIR=$(mktemp -d)

mkdir $VOLDIR/var

docker run -it --rm \
  -v $(pwd)/metadata:/etc/kaigara/metadata \
  -v $VOLDIR/var:/var/lib/testcase \
  kaigara/test:0.0.1

echo "After you can delete folder: ${VOLDIR}"
