#!/bin/sh

mkdir -p /var/lib/testcase

kaigara render server.conf > /var/lib/testcase/server.conf
