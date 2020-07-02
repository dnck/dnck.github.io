---
layout: post
title: List all tables in a postgresql database from python
author: dc
date: 2019-05-27
comments: true
analytics: true
keywords:  postgresql, python
description:
show_excerpt: true
excerpt: How to use the psycopg2 module in python to interact with a postgres server.
category: blog
---

```python
"""
List all tables in a postgresql database from python
"""
import psycopg2 as pg
import pandas.io.sql as sqlio

pd.set_option('display.float_format', lambda x: '%.4f' %x)
def do(cmd, conn):
    dat = sqlio.read_sql_query(cmd, conn)
    return dat
dname = ""
pguser = ""
passw = ""
conn = pg.connect("dbname={} user={} password={}".format(dname, pguser, passw))
cmd = """SELECT table_name FROM information_schema.tables
       WHERE table_schema = 'public'"""
do(cmd, conn)
```
