#!/usr/bin/env bash

awk 'BEGIN { FS=RS }
{ print length, $0 }' $* |
    sort -k1 |
    sed 's/^[0-9][0-9]*//'
