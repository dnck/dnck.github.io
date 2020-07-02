---
layout: post
title: udp server & client in python
author: dc
date: 2019-05-01
comments: true
analytics: true
keywords:  
description:
show_excerpt: true
excerpt: How to implement a basic udp client and server in python to measure response latency.
category: blog
---

I think a great way to learn about networks is to actually connect two computers
across an ocean. To that end, we'll use an EC2 Linux instance to send and receive
UDP datagrams. We will bake some logic into our code to measure the latency of responses,
and a primitive security feature for UDP clients.

To follow along with this tutorial, I assume the following:

* you already know how to launch an AWS ec2 instance
* you have python installed on your local machine
* you're operating on a Linux machine (although macOS might work as well)

## Step one: launch an ec2 instance & configure its inbound rules
Launch an ec2 Linux instance on aws and configure the inbound traffic rules
to accept only traffic from your home public IPv4.

You can get your public IPv4 via curl like so: ```curl ifconfig.me```.

For the rest of the tutorial, pretend your public IPv4 is ```XX.YYY.ZZZ.ABC```.

After you have your IPv4, you can change your ec2 inbound traffic rules to read
as follows:

| Port |  Protocol | IP (CIDR Notation)   |   
| ---- | ----| ----| ----|
| 80 |tcp	| XX.YYY.ZZZ.0/24
| 0-65535	| tcp	| XX.YYY.ZZZ.0/24
| 5432 |	tcp	| XX.YYY.ZZZ.0/24
| 22	| tcp	| XX.YYY.ZZZ.0/24
| 0-65535	| udp |	XX.YYY.ZZZ.0/24
| 443	| tcp |	XX.YYY.ZZZ.0/24

Replacing the final digits of your IPv4 (e.g .ABC) with .0/24 will allow
traffic from any computer on your home router/gateway
(unless you have some special config). So, if you don't trust the other computers
on your local net, consider change this (and changing house-mates!). Note that the
the integer after the / (24) specifies the range of addresses on
your local network that are permitted.
You can read more about that [here](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

After you set the traffic rules, ssh into your remote machine.

## Step two: provision your ec2 instance

On the remote machine, set up your minimal development environment like so,

```
sudo yum install python36.x86_64
curl -O https://bootstrap.pypa.io/get-pip.py
python3 get-pip.py --user
sudo yum install git
```

This will install Python 3.6, pip3, and git.

Now, you're ready to see the server-client code:

### UDP server-client example code

```python
#!/usr/bin/env python3
import argparse
import json
import random
import socket

from datetime import datetime

MAX_BYTES = 65535
SAVE_RESULTS = True
FAULTY_SERVER = False

def randomVal(upperbound=1000):
    return random.randint(0, upperbound)

def server(host, port, log=False):
    '''
    If using the remote, the the interface should correspond to
    the ip address of the machine.
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
        print('Request time: {}'.format(client_time))
    client_request = [request_ID, client_time]
    send_bytes = json.dumps(client_request).encode()
    sock.sendto(send_bytes, (host, port))
    received_bytes, address = sock.recvfrom(MAX_BYTES)
    server_reply = json.loads(received_bytes.decode())
    reply_ID = server_reply[0]
    if request_ID == reply_ID:
        print(server_reply[1])
        if log:
            print('Server time: {}'.format(server_reply[1]))
            print('Receipt time: {}'.format(datetime.now()))
    else:
        print('Server is a liar!')

if __name__ == '__main__':
    choices = {'client': client, 'server': server}
    parser = argparse.ArgumentParser(description=\
      'UDP server/client example with request ids and latency stats. \n'+\
      'Usage:\n '+\
      ' server: "python udp_remote_example.py server IPV4 PORT -measure_delay 1"'+\
      ' client: "python udp_remote_example.py client IPV4 PORT -measure_delay 1"'
    )
    parser.add_argument('role', choices=choices,
      help='which role to play: client or server'
    )
    parser.add_argument('host',  metavar='host', type=str, default='127.0.0.1',
      help='interface the server listens at, or host the client sends to'
    )
    parser.add_argument('port', metavar='port', type=int, default=1060,  
      help='UDP port (default 1060)'
    )
    parser.add_argument('--measure_delay', metavar='measure_delay',
      type=lambda s: s.lower() in ['true', 't', 'yes', '1'],
      default=False, help='measure the delay?'
    )
    args = parser.parse_args()
    function = choices[args.role]
    if args.measure_delay:
        if args.role == 'client':
            for x in range(1000):
                function(args.host, args.port, log=log)
        else:
            function(args.host, args.port, log=log)
    else:
        function(args.host, args.port)
```

### How to use
Save the above code somewhere on your local and remote machine. For the rest of
the tutorial, I assume you name the file, `cs-example.py`.

On the remote machine (ec2):
```
python3 cs-example.py server ec2-public-ip 8080 --measure_delay 1 --measure_delay 1
```

On the local machine:
```
python3 cs-example.py client ec2-public-ip 8080
```

After running the code, on the server, you should see something like this,
```
Server listening at ('123.4.5.678', 1234)
```

And once your client starts up, on the server, you should see something like this,
```
Request ID: 491. Client request, The time is 2019-05-04 16:20:56.975068
```

On the local machine, once the server responds, you should see something like,
```
The time is 2019-05-04 14:20:57.035891
```

which is the time according to the server.


## Review of UDP server-client example

The server above is just a time server. When a client sends its time to the server,
the server responds with its own time, which is the time at which the server
processes the client request. You can use this logic for all sorts of things.

Note that a simple security mechanism was hacked into the server. Do you see it?
To simulate a bad server, set ```FAULTY_SERVER``` to ```True```.
With this global set to True, the server will occasionally send over a response
with a secret id set by the client that doesn't match. In this way, the client
can know that the server is experiencing some difficulty in processing the request.

However, and this is important to note, this security mechanism is not really
secure. The client is still completely open to a man-in-the-middle attack, which
can be prevented by using a TLS wrapped socket. Maybe if someone requests me to
write about that, I will do a blog post.

Also note that if you set log to ```True``` and ```--measure_delay``` to ```True```,
then your home computer will print latency stats. In this way, you can measure
how long it takes to send/receive a response to your ec2 server.

Have fun with this! It was meant as a toy for your learning pleasure :)
