#!/bin/bash

if [ -z "$1" ]; then
  echo -e "\nERROR: Please specify new service name:\n\n\t./copy_template.sh <SERVICE_NAME>\n"
  exit 1
fi

export SERVICE_NAME=$1

export SERVICE_NAME_UPPER_FIRST=$(python -c "print(\"$SERVICE_NAME\".capitalize())")

export BASE_PATH="examples/${SERVICE_NAME}"

# Create folder structure
cp -R examples/template $BASE_PATH
mv $BASE_PATH/cmd/template              $BASE_PATH/cmd/$SERVICE_NAME
mv $BASE_PATH/internal/templateserver   $BASE_PATH/internal/${SERVICE_NAME}server
mv $BASE_PATH/rpc/template              $BASE_PATH/rpc/$SERVICE_NAME

# Find items files to replace in

# template --> name
export REPLACE_WITH="${SERVICE_NAME}"
export FIND_THIS=template
ITEMS=$(rg -l --color never "$FIND_THIS" "$BASE_PATH")
IFS=$'\n'

for ITEM in $ITEMS; do 
    sed -i '' "s/$FIND_THIS/$REPLACE_WITH/g" "$ITEM"
done

# Template --> Name
export REPLACE_WITH="${SERVICE_NAME_UPPER_FIRST}"
export FIND_THIS=Template
ITEMS=$(rg -l --color never "$FIND_THIS" "$BASE_PATH")
IFS=$'\n'

for ITEM in $ITEMS; do 
    sed -i '' "s/$FIND_THIS/$REPLACE_WITH/g" "$ITEM"
done