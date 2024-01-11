---
layout: post
title: Code Syntax Highlights with CSS
author: dc
date: 2021-08-07
comments: true
analytics: true
keywords: code, syntax, palette
description: This posts plays around with the code syntax color palette of darcula
category: blog
show_excerpt: true
excerpt: I got a lot of respect for good design!
---


After struggling with code-syntax highlighting in css, I have a lot of
respect for the frontend-devs that create color palettes for programmers.

As you can see on this page, I'm more of a backend, brute-force dev! Haha.

Anyway, in this post, I'll play around with the code syntax color palette. I'm
using [the Darcula Color Palette](https://draculatheme.com/contribute) with my common languages to get something
that hopefully isn't horrible on the eyes.

Right now, I have no idea what css variables control the appearance of the code
in the <pre>\<code> blocks. So, as always, trial-and-error is the best way to
learn!

First comes Python,

```python
# -*- coding: utf-8 -*-
import foo


class Foo():
    def __init__(self):
    """single line docstring"""
        pass
    def func(self):
        pass  # inline comment


if __name__ == "__main__":
    """This is a multiline
    comment
    """
    # this is a comment
    Foo()
```

Then comes go,

```go
import "fmt"

// foo is a function
func foo() {
  // this is a comment
  fmt.Println("foo")  // inline comment
}

func main() {
   foo()
}

```

The comes c++,

```cpp
#include <iostream>

// TODO: Implement Foo
class Foo {
 public:
  Foo(const std::string& name); // inline comment
  const std::string& GetName() const;
 private:
  const std::string name_;
};

Foo::Foo(const std::string& name)
    : name_(name) {
}

const std::string& Foo::GetName() const { return name_; }

int main(int argc, char** argv) {
  Foo foo("foo");
  std::cout << foo.GetName() << std::endl;
  return 0;
}
```

And bash,

```bash
# this is a comment
FOO="bar"
function hello() {
  echo $FOO
}
```

And Docker,

```Dockerfile
# comment
FROM docker.io/prom/prometheus
COPY config/prometheus.yaml /etc/prometheus/
```

And a Makefile,

```cmake
.PHONY: help
help: ## The default task is help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

help: ## just say hello
	 echo "hello world!"
```
