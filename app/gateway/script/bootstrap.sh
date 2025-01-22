#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=gateway
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}