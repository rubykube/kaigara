#!/bin/sh

if [[ $# -lt 1 ]] || [[ "$1" == "-"* ]]; then
  echo Running kaigara $@
  exec kaigara $@
fi

exec "$@"
