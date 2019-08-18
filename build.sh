#!/bin/bash

build_dir=`pwd`



echo 'building service...'
cd ${build_dir}
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -tags etcd -ldflags "-s -w" -o rpcx_service
cd ..
