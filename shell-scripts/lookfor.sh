#!/usr/bin/env bash

where="${1}"
when="${2}"
shift; shift

# build egrep expression like (word1|word2|..._

for word; do
  case "${expr}" in
    "") expr="(${word}" ;;
    *) expr="${expr}|${word}" ;;
  esac 
done
expr="${expr})"

find ${where} -type f -mtime ${when} -print | xargs grep -E "${expr}"

