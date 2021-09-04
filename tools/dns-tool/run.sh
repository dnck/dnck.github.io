#!/bin/sh
#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
# This is a simple shell script to run the Go binary `dns-tool`
# as a background process of the container. The log is saved as server.log.
#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

set -e

#./dns-tool >> ./server.log
./dns-tool > ./server.log 2>&1 &