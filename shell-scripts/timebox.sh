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
  while [[ $count -lt 300 ]]; do
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

/Users/nikhilthomas/Applications/Chrome\ Apps.localized/Home\ -\ Netflix.app/Contents/MacOS/app_mode_loader &>/dev/null &
pid=$(jobs -p | tail -n 1)

slice
