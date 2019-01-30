---
layout: post
title: Memory consumption in python programs
date: 2018-05-07
author: dc
comments: true
analytics: false
keywords: numpy, scipy, arrays, memory, garbage collection
description:
tag: Modelling, multisensory integration, taste, psychophysics
category: PhD study
---

L.T. recently showed me how to think about **adjacency matrices** for graphs, how to build them on the fly with **numpy** arrays.

But, now I got a problem. In a nutshell, I think that Python's transparent handling of memory allocation ain't so good for what I'm doing. For example, according to  O.F., Python doesn't have something called a **Garbage Collector**, which AFAIK frees up memory after a variable gets deleted from the current workspace. Instead, Python frees the space, but saves only for new variables that fit into the size of the freed up region. So, like, if I got a 200mb array, delete it, then, that space would be available, but only for something that's 200mb in size. That's pretty shitty, right?

Here's my current bottleneck. Maybe someone out there in lalaland can help.

I'm working with really big tables of numpy arrays (e.g. n x n, where n = 100,000,000).

First, I allocate space by creating this big empty table. Afterwards, I run functions in a loop, which add values to the cells in the table.

Some functions are relatively simple. For example, one function just puts a random number in some cells of a row.

Other functions are a bit more complicated.

For example, one function, call it a "random walk function", does something like,

    1) Start at this_row in the database
    2) Check if an adjacent_row to this_row has some value in it, and if so, then
    3) Move to that adjacent_row, and
    4) Repeat 1-3 until there is a good reason to stop (e.g. check the row for some criterion value).

When I run the above logic, Python seems to eat up my computer's memory.

I tried running a memory_profiler in Python, and it confirmed my intuition. The random walk function uses up a lot of my computer's memory. But, I don't have a solution. I suspect I'm missing some critical piece of knowledge or skill that would help me here.

Here's a few suggestions I've gotten already:

* Use Tensorflow
* Use parallel programming
* Use distributed programming
* Make your code more efficient
* Use C
* Go for a walk and come back later

Problem is -- I dunno which path to pursue.

Got any suggestions?
