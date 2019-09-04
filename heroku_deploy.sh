#!/bin/bash

set -e # Abort script at first error
set -u # Disallow unset variables

IMAGE_NAME=$1

# Deploy to heroku.
heroku container:push --app "go-animal-api" ${IMAGE_NAME}/web
heroku container:release web --app "go-animal-api"
