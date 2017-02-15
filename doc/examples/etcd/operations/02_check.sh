#!/bin/sh

set -xe

curl -L $PUBLIC_IP:5001/v2/stats/leader
