#!/usr/bin/env bash

SSSFile=~/.sssFile

function precmd() {
  notify-send precmdfn 
}
function preexec() {
  notify-send preexecfn
}
