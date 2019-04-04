---
layout: post
title: Logical operators in Python
author: dc
date: 2019-04-03
comments: true
analytics: true
keywords: java, concurrency
description:
category: blog
---

This post is a work in progress. It serves mainly as a reference to be updated occasionally on Python. I may in the future update this to include the equivalents in Java.

**OR (Inclusive) := |**
```
True | True == True
False | True == True
True | False == True
False | False == False
```
**OR (Exclusive) := ^**
```
True ^ True == False
False ^ True == True
True ^ False == True
False ^ False == False
```
