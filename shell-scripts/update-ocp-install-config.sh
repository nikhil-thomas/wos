#!/usr/bin/env bash
[[ -z ${1} || -z ${2} ]] && echo -e "usage:\n\t ${0}" '<new-cluster-name>' '<path-to-install-config>' && exit 0

INSTALL_CONFIG=${2}
NEW_CLUSTER_NAME=${1}
sed 's/^[[:space:]]*name: nikthoma.*/  name: '${NEW_CLUSTER_NAME}'/' ${INSTALL_CONFIG}

