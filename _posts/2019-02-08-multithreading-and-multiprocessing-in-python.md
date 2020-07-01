---
layout: post
title: Processes vs Threads
author: dc
date: 2019-02-08
comments: true
analytics: true
mathjax: true
keywords:  python
description: Processes vs. Threads
category: python
---

**What is a process?**  
A process is basically a container that holds all of the stuff that a program needs to run, including, references to the system resources available to the process.

Each process has an address space, a list of memory locations from 0 to some maximum that the process can read from and write to. The address space is important because it contains the executable program, the program data, and a stack.

**What is a thread?**
A thread, on the other hand, operates within the context of a process. One process can
utilize multiple threads, which share the address space of the process. For
this reason, concurrent programming with threads can be tricky if the data
is not safe from concurrent read/writes. For example, if a list maintained by
the address space of a process is written and read from two threads at the same
time within the process, then its quite possible that the reads are writes get
mixed up.
