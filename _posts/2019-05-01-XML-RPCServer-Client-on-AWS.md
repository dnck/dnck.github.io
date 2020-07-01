---
layout: post
title: UDP Time Server & Client on AWS EC2
author: dc
date: 2019-05-01
comments: true
analytics: true
keywords:  
description:
category: blog
---

While I think using localhost is good enough for education, I think a great way to learn about networks is to actually connect two computers across an ocean. To do that, I wanted to test using an EC2 Linux instance to send UDP & TCP datagrams and also run an XMLRPC-server on the instance.

Here's the steps I took to get this experiment working.  

**1st step**
Create a new account with AWS if you do not already have one.

**2nd step**
Create a free-tier EC2 Linux box and log into the Dashboard.

**3rd step**
Set the inbound traffic rules on the EC2 to only accept messages/traffic from your home ip.

On a unix terminal, ```curl ifconfig.me``` will display your ip.

For me, running the command returns, ```XX.YYY.ZZZ.ABC```.

Afer you have your ip, you can change your EC2 inbound traffic rules to read as follows:

| Port |  Protocol | IP (CIDR Notation)   |   
| ---- | ----| ----| ----|
| 80 |tcp	| XX.YYY.ZZZ.0/24
| 0-65535	| tcp	| XX.YYY.ZZZ.0/24
| 5432 |	tcp	| XX.YYY.ZZZ.0/24
| 22	| tcp	| XX.YYY.ZZZ.0/24
| 0-65535	| udp |	XX.YYY.ZZZ.0/24
| 443	| tcp |	XX.YYY.ZZZ.0/24

Replacing the final digits of your IP (e.g .ABC) with .0 will allow traffic from any computer on your home router/gateway (I think?). The final bits of the IP after the 0 (e.g. /24) specifies a range of addresses on the homenet. You can read more [here](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

After you set the traffic rules, you can copy and paste the ```ssh -i "yourkey.pem" ec2-user@ip``` from the EC2 Dashbord "connect" button into your terminal.

This should get you into the remote machine.

**4th step**

For this experiment, you can set up your minimal development environment like so,

```
sudo yum install python36.x86_64
curl -O https://bootstrap.pypa.io/get-pip.py
python3 get-pip.py --user
sudo yum install git
```

This will install Python 3.6, pip3, and git.

Ok, for the fun stuff.

**5th step**
You can clone some code from my github, or use the blocks below to send UDP datagrams from your home machine to the EC2 server.

Before looking at that code, I will tell you what it does, and how to use it.

**What it does**
The server is basically just a time server. The client will send his time to the server as a kind of request, and the server will respond with its own time. If ```FAULTY_SERVER``` is set to ```True```, then the server will occasionally send over a reply that doesn't match the secret id that the client sends over with its request. If you set log to ```True``` by running the functions on terminal with the ```-measure_delay``` set to ```True```, then your home computer will print out some detailed stats on the time regarding the sending and receiving of the messages. In this way, you can measure the latency. I leave it to you to send the log results to file.

**How to use**
On the server (ec2):
```
python3 xmlrpc_server_with_http_auth.py --host ec2-public-ip --port 8080 --username user --password pass
```
On the home machine:
```
python3 xmprpc_client.py --host ec2-public-ip  --port 8080 --username user --password pass
```
**What it looks like**
On the server, you'll see an output that looks like this,
```
Server listening at ('172.31.2.46', 1060)
```
And once you send a UDP Datagram, you'll see something like this:
```
Request ID: 491. Client request, The time is 2019-05-04 16:20:56.975068
```
On the client, you'll just see:
```
The time is 2019-05-04 14:20:57.035891
```
which is the time according to the server.

## UDP TIME SERVER WITH CLIENT
```python
#!/usr/bin/env python3
import argparse, socket
from datetime import datetime
import random
import json
import logging
import timeDelayLogger

MAX_BYTES = 65535
SAVE_RESULTS = True
FAULTY_SERVER = False

def randomVal(upperbound=1000):
    return random.randint(0, upperbound)

def server(host, port, log=False):
    '''
    if using the remote, the the interface should correspond to the ip address of the machine.
    '''
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sock.bind((host, port))
    print('Server listening at {}'.format(sock.getsockname()))
    while True:
        received_bytes, address = sock.recvfrom(MAX_BYTES)
        client_request = json.loads(received_bytes.decode())
        request = client_request[1]
        request_ID = client_request[0]
        print('Request ID: {}. Client request, {}'.format(request_ID, request))
        server_time =  'The time is {}'.format(datetime.now())
        if FAULTY_SERVER:
            if randomVal() > 500:
                request_ID+=1
        server_reply = [request_ID, server_time]
        send_reply = json.dumps(server_reply).encode()
        sock.sendto(send_reply, address)

def client(host, port, log=False):
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    request_ID = randomVal()
    client_time = 'The time is {}'.format(datetime.now())
    if log:
        log.logger.info('Request time: {}'.format(client_time))
    client_request = [request_ID, client_time]
    send_bytes = json.dumps(client_request).encode()
    sock.sendto(send_bytes, (host, port))
    received_bytes, address = sock.recvfrom(MAX_BYTES)
    server_reply = json.loads(received_bytes.decode())
    reply_ID = server_reply[0]
    if request_ID == reply_ID:
        print(server_reply[1])
        if log:
            log.logger.info('Server time: {}'.format(server_reply[1]))
            log.logger.info('Receipt time: {}'.format(datetime.now()))
    else:
        print('Server is a lying bastard!')

if __name__ == '__main__':
    choices = {'client': client, 'server': server}
    parser = argparse.ArgumentParser(description='Send and receive UDP locally with request ids. Usage:\n "python udp_remote_example.py server 127.0.0.1 1060 -measure_delay 0" \n Open a new terminal and, "python udp_remote_example.py client 127.0.0.1 1061 -measure_delay 0"')

    parser.add_argument('role', choices=choices, help='which role to play: client or server')
    parser.add_argument('host',  metavar='host', type=str, default='127.0.0.1', help='interface the server listens at, or host the client sends to')
    parser.add_argument('port', metavar='port', type=int, default=1060,  help='UDP port (default 1060)')
    parser.add_argument('-measure_delay', metavar='measure_delay', type=bool, default=False, help='measure the delay?')

    args = parser.parse_args()
    function = choices[args.role]
    if args.measure_delay:
        log = timeDelayLogger.ioManager()
        if args.role == 'client':
            for x in range(1000):
                function(args.host, args.port, log=log)
        else:
            function(args.host, args.port, log=log)
    else:
        function(args.host, args.port)
```

Here's another example, of an XMLRPC server on the EC2 instance. Here, you can send some Remote Procedure Calls from your home client computer to the EC2 box and with a little bit of modification have the EC2 box execute your Python code.


## Minimal RPC Server for Aws Ec2
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

## Notes
To get the above going, just set host and port to the EC2 public host and port (8080).
