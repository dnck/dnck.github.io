#!/bin/sh
#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
# This is a simple shell script to run the Go binary `dns-tool`
# as a background process and dump the log to server.log.
#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

set -e

./dns-tool > ./server.log 2>&1 &