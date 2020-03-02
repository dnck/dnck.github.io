---
layout: post
title: booleans-in-python-argparse
author: dc
date: 2020-03-02
comments: true
analytics: true
keywords: python
description: Python modules I know and love: argparse
category: blog
---

The `argparse` module in the python standard library is useful for building
command line tools. It allows you to run a python script and pass in some name options. Here's an very condensed example,

```
# -*- coding: utf-8 -*-
"""Create a new server
This script prints the output passed to the script on the command line.
"""
import argparse
if __name__ == '__main__':
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

On limit that I've found with argparse is when it comes to passing booleans on
the command line. For example, passing ```--capitalize True``` in the above
program does not result in the correct output.

You can get around this by defining the variable as a lambda expression. Thus,
we replace the argument ```capitalize``` in the above program like so,

```
parser.add_argument("message",
    metavar="--capitalize",
    type=lambda s: s.lower() in ['true', 't', 'yes', '1'],
    default=False
    help="Print the message in all caps."
)
```

So, now, this will always parse True, true, T, t, yes, Yes, and 1 as the
Python boolean ```True```.

Happy coding!
