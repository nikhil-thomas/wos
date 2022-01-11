#!/usr/bin/env bash

function lock_screen() {
  lockAfter=${1:-2}
  sleep $lockAfter
  if [[ $OSTYPE == "darwin"* ]]; then
    osascript mac-lock.scpt
  elif [[ $OSTYPE == "linux"* ]]; then
    xdg-screensaver lock
  fi
}

function speak() {
  message="$*"
  if [[ $OSTYPE == "darwin"* ]]; then
    say -v Daniel "[[volm 0.5]] $message"
  elif [[ $OSTYPE == "linux"* ]]; then
    espeak -p 50 -s 120 -a 250 "$message"
  fi
}

function slice() {
  count=0
  local duration=${1:-240}
  while [[ $count -lt ${duration} ]]; do
    sleep 10
    count=$((count + 10))
    echo $count
    speak "slice time"
  done
  speak "${brand_prefix}. 59 seconds to the end of this slice"
  sleep 50
  speak "${brand_prefix}. 10 seconds left"
  sleep 10
  speak "${brand_prefix}. end of slice"
  kill -15 ${pid}
  lock_screen
}

case ${1} in

  youtube)
    app='/Users/nikhilthomas/Applications/Chrome\ Apps.localized/YouTube.app/Contents/MacOS/app_mode_loader'
    ;;

  netflix)
    app='/Users/nikhilthomas/Applications/Chrome\ Apps.localized/Home\ -\ Netflix.app/Contents/MacOS/app_mode_loader'
    ;;


  kindle|read)
    app='/Applications/Kindle.app/Contents/MacOS/Kindle'
    echo test
    ;;

#  PATTERN_N)
#    STATEMENTS
#    ;;

  *)
    app='/Applications/Kindle.app/Contents/MacOS/Kindle'
    ;;
esac

echo $app

bash -c "$app" &> /dev/null &

pid=$(jobs -p | tail -n 1)

duration=${2:-240}
slice $duration
