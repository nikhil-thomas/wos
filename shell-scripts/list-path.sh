#!/usr/bin/env bash

lpath() {
  oldIFS=$IFS
  IFS=":"
  for p in $PATH; do
  echo $p
  done
  IFS=$oldIFS
}

lpath

