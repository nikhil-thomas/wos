#!/usr/bin/env bash

# Download Gnome extensions
# https://extensions.gnome.org/

# usage
# gnome-ext.sh <path-to-downloaded-extension>

EXT_PATH=$1
echo $EXT_PATH

EXT_UUID=$(unzip -c ${EXT_PATH} metadata.json | grep uuid | cut -d \" -f 4)
echo $EXT_UUID

GNOME_EXT_PATH=~/.local/share/gnome-shell/extensions/${EXT_UUID}

echo $GNOME_EXT_PATH
mkdir -p ${GNOME_EXT_PATH}

unzip -q ${EXT_PATH} -d $GNOME_EXT_PATH

#gnome-shell-extension-tool -e ${EXT_UUID}A
gnome-extensions -e ${EXT_UUID}A

