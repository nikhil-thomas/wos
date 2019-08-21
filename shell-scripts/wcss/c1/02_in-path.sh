#!/usr/bin/env bash

inPath() {
  local cmd=$1
  local givenPath=$2
  local oldIFS=$IFS
  local IFS=:
  local result=1

  for dir in $givenPath; do
    if [ -x $dir/$cmd ]; then
      result=0
      break
    fi
  done
  IFS=$oldIFS
  return $result
}

commandInPath() {
  local cmd=$1
  if [ $1 != "" ]; then
    if [ ${cmd:0:1} = "/" ]; then
      if [ ! -x  $cmd ]; then
        return 1
      fi
    elif ! inPath $cmd "$PATH"; then
      return 2
    fi
  fi
}

if [ $# -ne 1 ]; then
  echo "usage $0 <command>"
  exit 1
fi

commandInPath $1
case $? in
  0 )
      echo $1 found in PATH
      ;;
  1 )
    echo $1 not found or executable
    ;;
  2 )
    echo $1 not found in PATH
    ;;
esac
