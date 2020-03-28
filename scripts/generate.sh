#!/bin/bash

if [ -z "$1" ]; then
  echo -e "\nERROR: Please specify service name:\n\n\t./generate.sh <SERVICE_NAME>\n"
  exit 1
fi

export SERVICE_NAME=$1
export RPC_PATH=services/$SERVICE_NAME/rpc/$SERVICE_NAME/.

protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. $RPC_PATH/service.proto



