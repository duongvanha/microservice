#!/usr/bin/env bash

for d in ./../src/services/* ; do
  if [[ -f "$d/build.json" ]]; then
    echo "$d is a directory";
  fi;
done
