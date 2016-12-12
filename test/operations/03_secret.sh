#!/bin/bash

kaigara render secret.txt > /var/lib/testcase/secret.txt

kaigara render database.yml > /var/lib/testcase/database.yml
