#!/usr/bin/env zsh

install_config_path=${1:-'not-provided'}
cluster_number=${2:-01}
echo input-install-config=${install_config_path}
echo cluster_number=${cluster_number}
if [[ ! -f ${install_config_path} ]]; then
  echo 'install-config not found'
  exit 1
fi

sed -E "s/(name: nikthoma).*/\1.${$(date +"%b"):l}$(date +'%d').c${cluster_number}.$(openshift-install version | head -n 1 | cut -d ' ' -f 2 | tr -d ' ')/"  ${install_config_path} | tee install-config.yaml | grep name

openshift-install create cluster

