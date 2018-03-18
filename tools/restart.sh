#!/bin/sh

export GIN_MODE=release
./tools/stop.sh
nohup ./bin/myblog 2>&1 > ./log/nohup.log &


