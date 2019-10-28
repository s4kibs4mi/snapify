#!/bin/bash

export GO111MODULE=on
export GOARCH="amd64"
export CGO_ENABLED=0

cmd=$1

binary="snapify"

if [ "$cmd" = "build" ]; then
  echo "Executing build command"
  go get .
  go build -v -o ${binary}
  exit
fi

if [ "$cmd" = "serve" ]; then
  echo "Executing serve command"
  ./${binary} serve --config_path ./
  exit
fi

if [ "$cmd" = "up" ]; then
  echo "Executing migration up command"
  ./${binary} migration up --config_path ./
  exit
fi

if [ "$cmd" = "down" ]; then
  echo "Executing migration down command"
  ./${binary} migration down --config_path ./
  exit
fi

echo "No command specified"
