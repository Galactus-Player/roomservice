#!/bin/bash

export GO_POST_PROCESS_FILE=goimports
export OPENAPI_YML=./api/openapi.yaml

openapi-generator-cli generate --enable-post-process-file -i $OPENAPI_YML -g go-server
