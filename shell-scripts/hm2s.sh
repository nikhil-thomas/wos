#!/usr/bin/env bash

if [[ "$#" -lt 1 ]]; then
  echo "usage:"
  echo "       h2m  <hours> <minutes>"
  exit 1
fi

hrs=$1
mins=$2
secs=$(( $hrs * 60 * 60 + $mins * 60 ))
echo $secs

