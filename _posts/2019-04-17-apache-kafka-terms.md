---
layout: post
title: Apache Kafka - Chapter 1
author: dc
date: 2019-04-17
comments: true
analytics: true
keywords:  
description:
category: blog
---

> Every byte of data has a story to tell, something of importance that will inform the next thing to be done

> How we move the data becomes nearly as important as the data itself

## Publisher
A classifier of data, a creator of new messages to a specific topic

## Consumers
A consumer of classified data, a reader of messages, reads in the order of production, tracks what it consumes by tracking the offset, belong to a consumer group which works together, consumers consume an entire topic, a consumer is an owner of a partition from a publisher group

## Broker
A central point where messages are published, a single Kafka server, receives messages from producers, assigns offsets to them, commits the messages to storage on disk, services consumers, operate as part of a cluster, one broker in a cluster is a controller, a partition is owned by a single broker in a cluster - the leader

## Message
An array of bytes, the basic unit of data

## Offset
A bit of metadata, an integer value, monotonically increases

## Batches
A collection of messages written to the same partition / topic  

## Key
Metadata attached to a message. It's also a byte array


## Starting Zookeeper
Download zookeeper
move it to usr/local/zookeeper
```

$ tar -zxf zookeeper-3.4.6.tar.gz
$ mv zookeeper-3.4.6 /usr/local/zookeeper
$ mkdir -p /var/lib/zookeeper

$ cat > /usr/local/zookeeper/conf/zoo.cfg << EOF
> tickTime=2000
> dataDir=/var/lib/zookeeper
> clientPort=2181
> EOF

$ sudo /usr/local/zookeeper/bin/zkServer.sh start
JMX enabled by default
Using config: /usr/local/zookeeper/bin/../conf/zoo.cfg Starting zookeeper ... STARTED

$ sudo /usr/local/zookeeper/bin/zkServer.sh stop
```
On mac,
```
$ nc localhost 2181
$ srvr

Zookeeper version: 3.4.14-4c25d480e66aadd371de8bd2fd8da255ac140bcf, built on 03/06/2019 16:18 GMT
Latency min/avg/max: 0/0/0
Received: 1
Sent: 0
Connections: 1
Outstanding: 0
Zxid: 0x0
Mode: standalone
Node count: 4

```
