#!/usr/bin/env bash

i=0
echo before while
while (( i<5 ));
do
  echo $i
  i=$(( i + 1 ))
done
echo after while
echo ##########

echo before while
i=0
touch testfile
ls testfile

#while ls testfile;
#while ls testfile &> /dev/null;
while ls testfile > /dev/null 2>&1;
do
  echo $i
  sleep 2
  i=$(( i + 1 ))
  if [[ $i -eq 3 ]];
  then
    rm testfile
  fi
done
echo after while

