#!/bin/bash
NAME=$1
VERSION=$2
DOCKERFILE_DIR=$3
echo "Project: $NAME, Version: $VERSION, Path: $DOCKERFILE_DIR"

if [ $# -lt 3 ]; then
  echo "usage: bash docker_build.sh <PROJECT_NAME> <VERSION>"
  exit 1
fi
echo "Command: docker build -t simframe/$NAME:$VERSION $DOCKERFILE_DIR"

docker build -t simframe/$NAME:$VERSION $DOCKERFILE_DIR