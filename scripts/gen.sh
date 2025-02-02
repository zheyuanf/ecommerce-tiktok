#!/bin/bash

svcName=${1}

cd app/${svcName}
cwgo client -I ../../idl --type RPC --service ${svcName} --module github.com/zheyuanf/ecommerce-tiktok/app/${svcName} --idl ../../idl/${svcName}.proto
cwgo server -I ../../idl --type RPC --service ${svcName} --module github.com/zheyuanf/ecommerce-tiktok/app/${svcName} --idl ../../idl/${svcName}.proto
