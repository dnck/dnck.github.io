---
layout: post
title: xmlrpc server & client in python
author: dc
date: 2019-05-01
comments: true
analytics: true
keywords:  
description:
show_excerpt: true
excerpt: How to implement an xmlrpc server and client in python to execute your code remotely.
category: blog
---



With xmlrpc, you can use your local machine to make Remote Procedure Calls to
your remote machine. In this way, you can execute Python code remotely!

To follow along with this tutorial, I assume the following:

* you already know how to launch an AWS ec2 instance
* you have python installed on your local machine
* you're operating on a Linux machine (although macOS might work as well)
* you followed the instructions to from the previous post to secure the ec2 instance.

Let's just jump to the code,

## Minimal RPC Server
```python
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

## Minimal RPC Client for Aws Ec2
```python
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
## How to use
I assume you copy and pasted the code above into your ec2 instance and named
the file as, xmlrpc-server.py. On your local machine, you named the
client code, xmlrpc-client.py.

On the server,

```
python3 xmlrpc-server.py --host IPv4 --port 8080 --username user --password pass
```

On the client,

```
python3 xmlrpc-client.py  --host IPv4 --port 8080  --username user --password pass
```

After running the client, you should see the `msg` echoed back to you in uppercase.

Feel free to play around with this example. You can, for example, change the
echo function that runs on the server to something more complex. I've successfully
used this code to run simulation code on servers.
