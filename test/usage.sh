#!/bin/bash

set -x

export KAIGARA_METADATA=$(mktemp -d)
./kaigara
rmdir $KAIGARA_METADATA
