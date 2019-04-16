---
layout: post
title: Null bytes in Python
author: dc
date: 2019-04-16
comments: true
analytics: true
keywords:  
description:
category: blog
---
We're doing a lot of hashing and validation on our databases lately and I fooled myself into thinking that a **byte** representation of a **str** of 0s would be the same thing as a nullbyte * the number of zeroes.

But remember folks, **a byte representation of a str of 64 zeroes:**
```
b'0'*64
```
**does not make 64 null bytes:**
```
b'\x00'*64
```
*Let that be a lesson to you!*

### Code Snippet 
```
# *-* coding: utf-8 -*-
import hashlib
nullbytes768 = b'\x00'*768
print(hashlib.sha3_256(nullbytes768).hexdigest())
```
