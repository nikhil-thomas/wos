#!/usr/bin/env bash

echo $*
echo ---
oldIFS=$IFS
IFS=:
for x in $*; do
  echo $x
done
IFS=$oldIFS

