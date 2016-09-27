#!/bin/sh

sleep 3
set -xe

echo "Check kaigara fetching remote configuration: "
echo
kaigara render my.cnf
echo
