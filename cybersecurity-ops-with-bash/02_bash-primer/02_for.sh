#!/usr/bin/env bash

# for loop three variations

# 1. like for loops in other programming languages
# 2. loop through a list
# 3. loop through list of arbitary things

# 1.
for ((i=0; i<5; i++)); do
  echo loop t1 : i=$i
done

# 2.
echo "ARG: :$*:"
for ARG; do
  echo $ARG
done
echo "END"

# any variable can be used in place of ARG
# in the following example list = ARG = $*
for list; do
  echo $list
done
echo "END"

# 3.
echo ""
echo items
for item in a b  abc 1 3; do
  echo $item
done
echo done items

echo ""
echo list
list1=(1 b c 45 akf)
echo ${list1[*]}
for item in ${list1[*]}; do
  echo $item
done
echo end

echo ""
list2=(t h e 1 2 3 0 )
echo ${list[*]}
for x in ${list2[*]}; do
  echo $x
done
