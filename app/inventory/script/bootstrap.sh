#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/inventory"
exec "$CURDIR/bin/inventory"
