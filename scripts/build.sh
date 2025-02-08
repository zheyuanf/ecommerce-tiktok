#!/bin/bash

. scripts/list_app.sh

get_app_list

readonly root_path=`pwd`
for app_path in ${app_list[*]}; do
    svc_name=`basename ${app_path}`
    echo "Building ${svc_name}..."
    cd "${root_path}/${app_path}" &&  GOOS=linux go build -v -o "${root_path}/bin/${svc_name}/server"
done