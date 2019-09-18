#/usr/bin/env bash

for ARG; do
  echo $ARG
done

printf "\n#####\n\n"

i=1
for arg in $*; do

  echo $i : $arg
  let i=i+1
done

