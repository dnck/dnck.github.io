---
layout: post
title: Timing attacks
author: dc
date: 2019-04-03
comments: true
analytics: true
keywords: python, java, concurrency,security
description:
category: blog
show_excerpt: true
excerpt: Be careful when you write server code! If it can be timed, it can possibly be attacked!
---

A nice blog post on [Timing attacks](https://codahale.com/a-lesson-in-timing-attacks/) that says,
>Every time you compare two values, ask yourself: what could someone do if they knew either of these values?

If the answer is, they could exploit my program, then you might want to make sure your code executes at a constant time. For example, if you're checking whether two strings are equal to one another, you might be tempted to do something like this,
```python
def checkEquality(pass1, pass2, pass_size = 64):
  for i in range(pass_size):
    if pass1[i] != pass2[i]:
      return False # passwords do not match
  return True # passwords match
```
The problem with the code above is explained in the blog. If you can't think of a reason not to use that code, then I encourage you to read it in full. But, if you get already that the code finishes in varying amounts of time for non-matches, and that this fact can be used to guess the match, then you can go on to look at the alternative method,
```python
def is_equal(a, b):
    if len(a) != len(b):
        return False

    result = 0
    for x, y in zip(a, b):
        result |= x ^ y
    return result == 0
```
Thanks to the original poster. I learned something, and I'll be using that improved code in my projects.
