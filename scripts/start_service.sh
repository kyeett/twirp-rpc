#!/bin/bash

if [ -z "$1" ]; then
  echo -e "\nERROR: Please specify new service name:\n\n\t./copy_template.sh <SERVICE_NAME>\n"
  exit 1
fi

export SERVICE_NAME=$1

go run examples/$SERVICE_NAME/cmd/$SERVICE_NAME/main.go