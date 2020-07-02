---
layout: post
title: booleans-in-python-argparse
author: dc
date: 2020-03-02
comments: true
analytics: true
keywords: python
description: Python modules I know and love -- argparse
category: blog
show_excerpt: true
excerpt: How to deal with boolean flags passed to the argparse module in python.
---

The `argparse` module in the Python standard library is useful for building
command line tools. One "gotcha", however, is parsing Boolean values from the user's
input.

Here's an very simple example,

```py
# -*- coding: utf-8 -*-
"""Create a new server
This script prints the output passed to the script on the command line.
"""
import argparse

if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="Demo argparse"
    )
    parser.add_argument("message",
        metavar="message",
        type=str,
        help="The message to print on the command line."
    )
    parser.add_argument("message",
        metavar="--capitalize",
        type=str,
        default=False
        help="Print the message in all caps."
    )
    args = parser.parse_args()
    if args.capitalize:
      print(args.message.upper())
    print(args.message)
```

Passing ```--capitalize True``` to the above program would produce an unexpected
output. That's because `argparse` doesn't handle booleans. One alternative would
be to use the excellent `click` module instead of `argparse`.

But, what if you're just writing something simple, and you want to get the job
done quickly?

Here's a nice Pythonic solution: define the type of your input parameter as a
lambda expression that parses the expected (and common variants) input.

Thus, to get a boolean parameter using the `argparse` module, simply make the type
of your argument like so,

```python
parser.add_argument("message",
    metavar="--capitalize",
    type=lambda s: s.lower() in ['true', 't', 'yes', '1'],
    default=False
    help="Print the message in all caps."
)
```

Now, this will parse the command line passed strings:

[ True, true, T, t, yes, Yes, 1]

as the Python boolean ```True```.

Hope this helps. Happy coding!
