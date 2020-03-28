#!/bin/bash

if [ -z "$1" ]; then
  echo -e "\nERROR: Please specify service name:\n\n\t./start_service.sh <SERVICE_NAME>\n"
  exit 1
fi

export SERVICE_NAME=$1

go run examples/$SERVICE_NAME/cmd/$SERVICE_NAME/main.go