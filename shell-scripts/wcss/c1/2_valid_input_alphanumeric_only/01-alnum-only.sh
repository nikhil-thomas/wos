#!/usr/bin/env bash

validAlphaNum() {
  # validate arg
  # return 0 if all chars are upper/lower/digits
  # retyrn 1 otherwise

  validChars="$(echo $1 | sed -e 's/[^[:alnum:]]//g')"
  echo validChars=$validChars

  if [[ $1 == $validChars ]]; then
    return 0
  else
    return 1
  fi
}

validUpperSpace() {
  # validate arg
  # return 0 if all chars are upper/space
  # retyrn 1 otherwise

  validChars="$(echo $1 | sed -e 's/[^[:upper:] ]//g')"
  echo validChars=$validChars

  if [[ $1 == $validChars ]]; then
    return 0
  else
    return 1
  fi
}

validPhoneNumber() {
# validate arg
# return 0 if all chars are digit, - , (), no leading spaces or multiple spaces
# retyrn 1 otherwise

validChars="$(echo $1 | sed -e 's/[^- [:digit:]\(\)]//g')"
echo validChars=$validChars

if [[ $1 == $validChars ]]; then
  return 0
else
  return 1
fi
}

#==================

echo -n "Enter ur input:"
read input
if ! validAlphaNum "$input"; then
  echo "Input is not valid alphanum"
  #exit 1
else
  echo "Input is valid alphanum"
fi

if ! validUpperSpace "$input"; then
  echo "Input is not valid upperSpace"
  #exit 1
else
  echo "Input is valid upperSpace"
fi
if ! validPhoneNumber "$input"; then
  echo "Input is not valid phone number"
  #exit 1
else
  echo "Input is valid phonenumber"
fi
exit 0
