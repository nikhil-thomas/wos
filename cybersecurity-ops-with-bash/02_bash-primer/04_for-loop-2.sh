#!/usr/bin/env bash

for item in $(ls | grep sh); do
  echo $item:
done

echo "#########"
for item in $(ls | grep sh) {1..5}; do
  echo $item
done

