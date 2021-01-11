#!/usr/bin/env bash

#OWNER_KIND="TektonConfig1212"
#OWNER_APIVERSION="operator.tekton.dev/v1alpha1123123"
#OWNER_NAME="config123123"
#OWNER_UID="111111123123"
#
#function findOwnerUID {
#    OWNER_UID=$(kubectl get tektonconfig config -o=jsonpath='{.metadata.uid}')
#    #[[ -z $OWNR_UID ]] || return 1 || return 0
#}

function checkOwner() {
    ct=$1
    configOwner=$(kubectl get clustertasks $ct -o=jsonpath='{.metadata.ownerReferences[?(@.name=="cluster")].kind}')
    if [[ -n $configOwner ]]; then
	return 0
    fi
    return 1
}

#function updateOwnerReference() {
#    ct=$1
#    kubectl patch clustertask $ct --type='json' -p='[{"op": "remove", "path": "/metadata/ownerReferences"}]'
#
#    kubectl patch clustertask $ct --type='json' -p=\[\{\"op\":\ \"add\",\ \"path\":\ \"/metadata/ownerReferences\",\ \"value\":\ \[\{\"name\":\"$OWNER_NAME\",\ \"kind\":\"$OWNER_KIND\",\ \"apiVersion\":\ \"$OWNER_APIVERSION\",\ \"uid\":\ \ \"$OWNER_UID\"\}\]\ \}\]
#}

#findOwnerUID

clustertasks=$(kubectl get clustertasks -o=jsonpath='{.items[*].metadata.name}')

for ct in $clustertasks; do
    echo $ct
    checkOwner $ct
    echo $?
    if checkOwner $ct; then
	echo change owner
#	updateOwnerReference $ct
    fi
    
done
