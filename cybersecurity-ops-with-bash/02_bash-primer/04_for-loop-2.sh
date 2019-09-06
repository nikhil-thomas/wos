#!/usr/bin/env bash

for item in $(ls | grep sh); do
  echo $item:
done

echo "#########"
for item in $(ls | grep sh) {1..5}; do
  echo $item
done

echo "#########"

for i in {99..106..2}; do
  echo $i
done

echo "#########"

for i in {099..106..2}; do
  echo $i
done
