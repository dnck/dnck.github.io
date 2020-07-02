---
layout: post
title: Network programming in python
author: dc
date: 2019-01-22
comments: true
analytics: true
keywords: urllib, python, client-server programs, networks, python, lesson
description: a short HTTPclient in python
category: blog
show_excerpt: true
excerpt: Download a webpage using python urlib module.
---

The urllib is really powerful. It can wrap your input data in headers and send them off to a server as properly formatted html request objects at a remote host. Saying this entails that we distinguish between the concept of a host and a server. Whereas a host is the suite of protocols that define the goals of a target operating system, a server is a program that runs on a host. Many people do not make this distinction. They do not realize that their pcs and smartphones can be configured to run as server apps. This should not be surprising, given the news about hacking and all, but it may seem surprising to some. For those of you who are concerned that your computer is participating in computations not under your control, I have to say, I sympathize with you, but *I DO plan to inspect all of the code that runs on my pc all of the time!* It's just that sometimes ... deadlines.

Ok, now check out some cool program code:

```python
#!/usr/bin/env python
# Python 3.7
# download_data.py

import argparse
import urllib.request

REMOTE_SERVER_HOST = 'http://www.cnn.com'

class HTTPClient:
    def __init__(self, host):
        self.host = host

    def fetch(self):
        response = urllib.request.urlopen(self.host)

        data = response.read()
        text = data.decode('utf-8')
        return text

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='HTTP Client Example')
    parser.add_argument('--host', action="store", dest="host",  default=REMOTE_SERVER_HOST)
    given_args = parser.parse_args()
    host = given_args.host
    client = HTTPClient(host)
    print (client.fetch())
```

This program basically does something like,

  Step 1. Fetch some utf-8 symbols from the host, http://www.cnn.com.
  Step 2. Decode the utf-8 symbols and print them in ascii characters on the standard output

Try running it for yourself.

It will give you some interesting stuff to read in the terminal. You can also try it with some parameters.

It works like this,

```
python download_data.py --host=http://www.danjcook.com
```
The program constructs an HTTPClient class. An instance of the class assigns to itself a name corresponding to an ip address of a remote host, REMOTE_SERVER_HOST. The client then has a single method that does all of the things we mentioned at the outset of the lesson. It constructs http headers, and sends them off to a server as properly formatted requests for html at the remote host. On the receiving end, it receives replies from the server of the data. Note that state changes triggered in the server app running on the host may have resulted in a complex interplay of several different peices of technology, but we do not know this. It is transparent to us.

Thus, we will have to follow up in the future by showing how to request from the remote host their credentials. We can then use this information to make decisions about the soundness of the response.
