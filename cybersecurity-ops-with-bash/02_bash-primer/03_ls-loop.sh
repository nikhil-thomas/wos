#!/usr/bin/env bash

files=($(ls))
echo ${files[*]}

for file in ${files[*]}; do
  echo $file
done

