---
layout: post
title: PostGres Python API
author: dc
date: 2019-04-13
comments: true
analytics: true
keywords:  
description:
category: blog
---
Here's are some commands you might find useful for dealing with PostgreSQL and interacting with it via the python api psycopg2:

## PostgreSQL
**Installation of postgresql on ubuntu / linux with python bindings:**

```
apt-get install postgresql-10
pip install psycopg2-binary
```

**Configure postgresql for linux-like distributions including macos:**
* login to the default postgres username with ```sudo su - postgres```, then enter ```psql```

On mac, you need to export the /bin directory where the ```psql``` command lives when you log in as the postgres user. You can find this in the pg installation notes. For me, this means I do this every time I login to the postgres user: ```export PATH=/Library/PostgreSQL/11/bin:$PATH```

* Remember, each new project needs its own database. So create it. 

* Once logged into postgres, create the database for the project with ```CREATE DATABASE projectname;```

* Create a user with ```CREATE USER user WITH PASSWORD 'password';```

  * Yes, you need the quotes around the password you enter.
  * You, remembert to add the ; to the end of the psql command. 

You can also set up some things with the database for speed and optimization. For example,

* ```ALTER ROLE user SET client_encoding TO 'utf8';```

* ```ALTER ROLE user SET default_transaction_isolation TO 'read committed';```

* ```ALTER ROLE user SET timezone TO 'UTC';```

If you don't know what they mean, you should probably take a step back and read a bit about distributed systems, esp. concurrency control. I recommend the book, [Distributed Systems by author A. Tanenbaum](https://www.distributed-systems.net/index.php/books/distributed-systems-3rd-edition-2017/).

* Give the username all access on the db: ```GRANT ALL PRIVILEGES ON DATABASE projectname TO user;```

* Exit psql with ```\q``` and then ```exit```

## Python, psycopg2
Let's do a run through of some basic commands using the python api. 

Make sure you did, ```pip install psycopg2-binary```. 

Then, 

```
import psycopg2 as pg

# replace below w info you created with psql 
dname = "template0"
pguser = "username"
pass = "password"

# connect to the db:

conn = psycopg2.connect("dbname={} user={} password={}".format(dname, pguser, pass))

# get a cursor to interact w db:

cur = conn.cursor()

# view all tables in db:

def print_all_tables(cur):
  tables = []
  cur.execute("""SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'""")
  for table in cur.fetchall():
    tables.append(table)
  return tables

# use of function defined above:

tables = print_all_tables(cur)

# check out tables if you like:

print(tables) # its a list of tuples :) easy!


# get the info from the 0 table as py objects:

cur.execute("SELECT * FROM {};".format(tables[0][0]))
cur.fetchall()

# careful with this one. It prints all table names and their infos:

def print_all_tablesinfo(tables):
  for tablename in tables:
    tablename = tablename[0]
    print(tablename, "\n")
    cur.execute("SELECT * FROM {};".format(tablename))
    print(cur.fetchall())

# print_all_tablesinfo(tables)

```

And that's that.

One important take away is that the stuff you return from postgres to your interpreter could possibly be python objects, e.g. like a ```datetime.datetime``` class. If so, then you'll be able to call ordinary methods of the class plus get any attributes of the class in the ordinary Pythonic way! That means, postgres is pretty powerful indeed :)

Last thing: you might also be interested in the Django tutorials which go through creating a basic webapp which uses postgres. I think there's something on my github with the walk through notes, but I couldn't be bothered to link you to it. Sorry ;) 
