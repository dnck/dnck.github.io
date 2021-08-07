---
layout: post
title: code-syntax-highlighting
author: dc
date: 2021-08-07
comments: true
analytics: true
keywords:
description:
category: blog
show_excerpt: true
excerpt: I got a lot of respect for good design!
---


After failing to get nice looking code-syntax highlighting, I have a lot of
respect for the frontend-devs that create beautiful designs with colors.
As you can see on this page, I'm more of a backend, brute-force dev. Haha.

Anyway, I'm playing around with some syntax highlighting. You can see my
colors and rules in the _sass/iiibit directory. It's a bit of a mess right
now, so I'm going to use this page to experiment with different languages
to see what css variables control the appearance of the code in the <pre>\<code> blocks.

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


if __name__ == '__main__':
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

```c++
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

And finally, bash,

```bash
# this is a comment
FOO="bar"
function hello() {
  echo $FOO
}
```

Oddly, as you can see, my current styles need to be adjusted. I thought I was
actually getting orange fo comments in python, but look at that, I got them in
every language except python!! And tbh, I'm feeling rusty on Python. I'm not
even sure if that py program would run!

