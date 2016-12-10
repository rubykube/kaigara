#!/bin/bash

assert()
{
  E_ASSERT_FAILED=72
  E_PARAM_ERR=98

  if [ -z "$2" ]
  then
    return $E_PARAM_ERR
  fi

  lineno=$2

  if [ ! $1 ]
  then
    echo "Assertion failed:  \"$1\"" >&2
    echo "File \"$0\", line $lineno" >&2
    exit $E_ASSERT_FAILED
  fi  
} 

assert "-f /etc/kaigara/metadata/server.yaml" $LINENO

assert "-f /etc/kaigara/metadata/values.yml" $LINENO

assert "-f /var/lib/testcase/check" $LINENO

code=$(cat /var/lib/testcase/check)
assert "'$code' = 'fGkpkvT97sSe3dnc'" $LINENO

assert "-f /var/lib/testcase/secret.txt" $LINENO

secret=$(cat /var/lib/testcase/secret.txt)
assert "'$secret' = 'TqM4Fx7BZU9WE1iE6aRBu5JvL6dNrN5v'" $LINENO
echo "All tests pass"
