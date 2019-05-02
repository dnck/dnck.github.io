---
layout: post
title: How to XMLRPC Server & Client on AWS
author: dc
date: 2017-05-01
comments: true
analytics: true
keywords:  
description:
category: blog
---

This is going to be really quick. I was talking to a colleague at work yesterday, and I suggested he create a AWS EC2 Linux box and get it working for a Java Restful API + UDP/TCP Server Fullnode we're working on. Anyway, I wanted to make sure this task wasn't some impossible mission. So, tonight, I logged a new account with AWS, created a free-tier EC2 Linux box, set my inbound traffic rules to only accept messages/traffic from my home net, and bam! I couldn't get UDP messages passed from my home client machine to the EC2 server machine. However, I could run an XMLRPC server on the linux box, and send some Remote Procedure Calls from my home client computer to the EC2 box. So, without further ado, here's the code for doing that on server and client:


## Server
```
import argparse
import xmlrpc
from base64 import b64decode
from xmlrpc.server import SimpleXMLRPCServer, SimpleXMLRPCRequestHandler
# Python 3.6

class SecureXMLRPCServer(SimpleXMLRPCServer):

    def __init__(self, host, port, username, password, *args, **kargs):
        self.username = username
        self.password = password
        class VerifyingRequestHandler(SimpleXMLRPCRequestHandler):
              def parse_request(request):
                  if SimpleXMLRPCRequestHandler.parse_request(request):
                      if self.authenticate(request.headers):
                          return True
                      else:
                          request.send_error(401, 'Authentication failed, Try agin.')
                  return False
        SimpleXMLRPCServer.__init__(self, (host, port), requestHandler=VerifyingRequestHandler, *args, **kargs)

    def authenticate(self, headers):
        headers = headers.get('Authorization').split()
        basic, encoded = headers[0], headers[1]
        if basic != 'Basic':
            print ('Only basic authentication supported')
            return False
        secret = b64decode(encoded).split(b':')
        username, password = secret[0].decode("utf-8"), secret[1].decode("utf-8")
        return True if (username == self.username and password == self.password) else False


def run_server(host, port, username, password):
    server = SecureXMLRPCServer(host, port, username, password)
    def echo(msg):
        reply = msg.upper()
        print ("Client said: %s. So we echo that in uppercase: %s" %(msg, reply))
        return reply
    server.register_function(echo, 'echo')
    print ("Running a HTTP auth enabled XMLRPC server on %s:%s..." %(host, port))
    server.serve_forever()

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Multithreaded multicall XMLRPC Server/Proxy')
    parser.add_argument('--host', action="store", dest="host", default='localhost')
    parser.add_argument('--port', action="store", dest="port", default=8000, type=int)
    parser.add_argument('--username', action="store", dest="username", default='user')
    parser.add_argument('--password', action="store", dest="password", default='pass')
    given_args = parser.parse_args()
    host, port =  given_args.host, given_args.port
    username, password = given_args.username, given_args.password
    run_server(host, port, username, password)
```

## Client
```
import argparse
import xmlrpc
# Python 3.6

def run_client(host, port, username, password):
    server = xmlrpc.client.ServerProxy('http://%s:%s@%s:%s' %(username, password, host, port, ))
    msg = "hello server..."
    print ("Sending message to server: %s  " %msg)
    print ("Got reply: %s" %server.echo(msg))

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Multithreaded multicall XMLRPC Server/Proxy')
    parser.add_argument('--host', action="store", dest="host", default='localhost')
    parser.add_argument('--port', action="store", dest="port", default=8000, type=int)
    parser.add_argument('--username', action="store", dest="username", default='user')
    parser.add_argument('--password', action="store", dest="password", default='pass')
    given_args = parser.parse_args()
    host, port =  given_args.host, given_args.port
    username, password = given_args.username, given_args.password
    run_client(host, port, username, password)
```

## Notes
To get the above going, just set host and port to the EC2 public host and port (8080).
