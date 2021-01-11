#!/usr/bin/env bash

sssFile=${HOME}/.sssFile
trap 'rm ${sssFile}' DEBUG

while true; do
  touch ${sssFile}
  sleep 5
  [[ -f ${sssFile} ]] && cmatrix -s
done&

