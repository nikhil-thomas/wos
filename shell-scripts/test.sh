#!/usr/bin/env bash

sleep 100 &
sleep 40&
/Users/nikhilthomas/Applications/Chrome\ Apps.localized/Home\ -\ Netflix.app/Contents/MacOS/app_mode_loader &>/dev/null &

pid=$(jobs -p | tail -n 1)
sleep 5
kill -15 $pid
echo end