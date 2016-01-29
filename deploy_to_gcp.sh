#!/bin/bash

set -e

KUBERNETES_APP_NAME=total-sudoku
KUBERNETES_RC_NAME=sudoku-api
GOOGLE_CONTAINER_NAME=gcr.io/$KUBERNETES_APP_NAME/$KUBERNETES_RC_NAME:$CI_BUILD_ID
DEFAULT_ZONE=us-central1-f

codeship_google authenticate

echo "Setting default timezone $DEFAULT_ZONE"
gcloud config set compute/zone $DEFAULT_ZONE

echo "Configure kubectl"
gcloud container clusters get-credentials $KUBERNETES_APP_NAME

echo "Building docker image"
docker build -t chrismar035/$KUBERNETES_RC_NAME /deploy

echo "Tagging the Docker machine for Google Container Registry push"
docker tag -f chrismar035/$KUBERNETES_RC_NAME $GOOGLE_CONTAINER_NAME

echo "Pushing to Google Container Registry: $GOOGLE_CONTAINER_NAME"
gcloud docker push $GOOGLE_CONTAINER_NAME

echo "Starting rolling update of $KUBERNETES_RC_NAME with image $GOOGLE_CONTAINER_NAME"
kubectl rolling-update $KUBERNETES_RC_NAME --image=$GOOGLE_CONTAINER_NAME
