#!/usr/bin/env bash
t=${1:-10s}
message=${2:-time out, switch to next task}

#go get github.com/antonmedv/countdown
countdown $1
espeak -s 150 -p 1 -v en-us "$message"

