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
show_excerpt: true
excerpt: The logical distinction between the OR operator in Python.
---


**OR (Inclusive) := |**
```python
True | True == True
False | True == True
True | False == True
False | False == False
```
**OR (Exclusive) := ^**
```python
True ^ True == False
False ^ True == True
True ^ False == True
False ^ False == False
```
