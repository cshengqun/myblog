#!/bin/sh
if [ "$1" = "darwin" ]
then
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/myblog
elif [ "$1" = "windows" ]
then
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/myblog
else
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/myblog
fi
