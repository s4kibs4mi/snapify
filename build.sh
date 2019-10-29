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
  ./${binary} serve --config_path ./ --config_name config.example
  exit
fi

if [ "$cmd" = "instant" ]; then
  echo "Executing instant command"
  ./${binary} instant --url "$2" --out "$3" --config_path ./ --config_name config.example
  exit
fi

if [ "$cmd" = "up" ]; then
  echo "Executing migration up command"
  ./${binary} migration up --config_path ./ --config_name config.example
  exit
fi

if [ "$cmd" = "down" ]; then
  echo "Executing migration down command"
  ./${binary} migration down --config_path ./ --config_name config.example
  exit
fi

if [ "$cmd" = "docker" ]; then
  echo "Executing docker build command"
  docker build -t s4kibs4mi/snapify:"$2" .
  exit
fi

echo "No command specified"
