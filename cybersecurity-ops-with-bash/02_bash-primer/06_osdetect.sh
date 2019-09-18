#/usr/bin/env bash

if type -t wevtutil &> /dev/null; then
  OS=MSWin
elif type -t sctutil &> /dev/null; then
  OS=macOS
else
  OS=Linux
fi
echo $OS

