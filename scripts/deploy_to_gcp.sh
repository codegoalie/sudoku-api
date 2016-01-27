#!/bin/bash

set -e

if [ -n $CI_BUILD_ID ]
then
  $CI_BUILD_ID=local-build
fi

KUBERNETES_APP_NAME=total-sudoku
KUBERNETES_RC_NAME=sudoku-api
GOOGLE_CONTAINER_NAME=gcr.io/$KUBERNETES_APP_NAME/$KUBERNETES_RC_NAME:$CI_BUILD_ID
DEFAULT_ZONE=us-central1-f

codeship_google authenticate

echo "Setting default timezone $DEFAULT_ZONE"
gcloud config set compute/zone $DEFAULT_ZONE

echo "Configure kubectl"
gcloud container clusters get-credentials $KUBERNETES_APP_NAME

echo "Starting rolling update of $KUBERNETES_RC_NAME with image $GOOGLE_CONTAINER_NAME"
kubectl rolling-update $KUBERNETES_RC_NAME --image=$GOOGLE_CONTAINER_NAME
