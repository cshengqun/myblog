#!/bin/sh

pID=`pgrep myblog`
kill -9 ${pID}
