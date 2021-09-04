#!/bin/sh
#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
# This is a simple shell script to run the Go binary `dns-tool`
# as a background process and dump the log to server.log.
# If, instead you want to run the Docker image, then use,
# ```
# make build &&\
# sudo docker run -d --dns 127.0.0.1 --name dns-tool dnck.github.io/dns-tool &&\
# sudo docker exec -it dns-tool dig +tcp www.danjcook.com &&\
# sudo docker logs dns-tool
# ```
# Thanks!
# - Daniel Cook
#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

set -e

./dns-tool > ./server.log 2>&1 &