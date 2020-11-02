---
layout: post
title: new-py-project-templates
author: dc
date: 2020-11-02
comments: true
analytics: true
keywords: python, templates, agile
description: A comprehensive set of instructions for automating a new Python project
category: blog
show_excerpt: true
excerpt: A comprehensive set of instructions for automating a new Python project
---

This blog post contains a set of files and bash scripts for creating a new python
project in a virtual environment. It installs a command line tool for the new
app, and defines a minimal Dockerfile, and Makefile, which can be used for cloud
deployments. It's my hope that you find this helpful, and it minimizes the time
you spend creating new projects that follow best practices.


## Step 1.

Create the following files.

```bash
touch $HOME/Templates/python_class.py
touch $HOME/Templates/python_cli.py
touch $HOME/Templates/python_main.py
touch $HOME/Templates/python_makefile
touch $HOME/Templates/python_setup.py
touch $HOME/Templates/python_test_context.py
touch $HOME/Templates/python_unittest.py
touch $HOME/Templates/python.dockerfile
```

We will go through the code of each of these files.


### python_setup.py
```python
import setuptools

with open("./README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="app", # Replace with your own project name
    version="0.0.1",
    author="Your Name",
    author_email="youremail@somedomain.com",
    description="This project is just great!",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://danjcook.com/git",
    packages=setuptools.find_packages(),
    # include_package_data=True,
    install_requires=[
        "click==7.1.1",
        "python-dotenv==0.11.0"
    ],
    entry_points="""
        [console_scripts]
        app=main:cli
    """,
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: MIT License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.7',
)
```

We are going to install a command line tool for this project, so that you can
pass parameters and arguments to the program easily. We will use the excellent
module, Click. We will also define our secrets in a .env file, and use the
python-dotenv module for securely accessing our secrets at runtime.

### python_class.py
```python
# -*- coding: utf-8 -*-
"""Module description

This module does great things.
"""
import os

# Implementation constants
SCRIPT_DIRNAME, SCRIPT_FILENAME = os.path.split(os.path.abspath(__file__))
PROJECT_ROOT_DIR = os.path.dirname(SCRIPT_DIRNAME)

# Classes, methods, functions, and variables
class AppName():
    def __init__(self):
        pass
    def func(self):
        pass
```

You will probably have a set of classes in your project that do different things.
This will help to keep your project organized and understandable. Here, we define
the first class for our project and create a new file for it.

### python_cli.py
```python
import os
import sys
import click

from app import app

class AppNameCLI(object):
    def __init__(self):
        self.app = app.AppName()

@click.group()
@click.pass_context
def cli(ctx, settings_file):
    ctx.obj = AppNameCLI()

@cli.command()
# @click.argument('arg1')
# @click.argument('arg2')
# @click.option('--dryrun', is_flag=True)
@click.pass_context
def app_func(ctx):#, arg1, arg2, dryrun):
    """Description:
    \b

    Hello world
    """
    server = ctx.obj.app.func()

if __name__ == '__main__':
    cli()
```

We will access our class that does the processing from the command line. Here,
we import and do the work.

### python_main.py
```python
# -*- coding: utf-8 -*-
"""Module description

This module does great things.

Example:
    $ python3 great_things.py positional_argument --keyword_argument 1

Style guide: https://www.python.org/dev/peps/pep-0008/
"""

import argparse
import os

from app import app

# Implementation constants go ...
# here

if __name__ == '__main__':

    parser = argparse.ArgumentParser(
        description=""
    )
    parser.add_argument("--optional",
        metavar="opt",
        type=str,
        default="Add an option",
        help="Add an optional argument"
    )

    args = parser.parse_args()

    app = app.AppName()
    app.func()
```

This code is for legacy that does not want to transition to the use of the Click
module. It is not necessary, but it can be helpful for quick one-off projects.

### python_makefile
```txt
# If you got podman installed, we're going to build with it!
PODMAN := $(shell command -v podman 2>/dev/null)
DOCKER := $(shell command -v docker 2>/dev/null)
ifeq ($(PODMAN), /usr/bin/podman)
CONTAINER_ENGINE := $(PODMAN)
else
CONTAINER_ENGINE := $(DOCKER)
endif

.PHONY: help
help: ## The default task is help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

build_image: help_deploy## Build the image
	 "$(CONTAINER_ENGINE)" build -f ./Dockerfile -t appname
	 "$(CONTAINER_ENGINE)" image prune --filter label=stage=appname_base_image

prune_baseimage: ## Remove the intermediate image
	 "$(CONTAINER_ENGINE)" image prune --filter label=stage=appname_base_image

run_test: ## Run app for a help message after image build
	 "$(CONTAINER_ENGINE)" run app

ship_image: ## Save image as tar, scp to and load the image on host machine
	"$(CONTAINER_ENGINE)" save > appname.tar appname
	scp appname.tar server:~/appname/
	ssh server "cd ~/appname && docker load -i appname.tar"

distribute: build_image ship_image ## build image, save tar, scp, load

help_deploy: ## Build image and run tests in container
	@echo "--------------------------------------------------------------------"
	@echo "Usage:"
	@echo "podman run \\"
	@echo "    appname <options> <commands> <options>"
	@echo "--------------------------------------------------------------------"
```

No project is complete without continuous integration and delivery built in.
This Makefile will define some simple concepts for building, testing, and
distributing your project.

### python_test_context.py
```python
# -*- coding: utf-8 -*-
#pylint: skip-file
"""
This script is a testing utility used to append the '../app' directory
(where the primary .py executables live) to the system path so that they can be
imported during the tests.
"""
import os
import sys


SCRIPT_DIRNAME, SCRIPT_FILENAME = os.path.split(os.path.abspath(__file__))
PROJECT_ROOT_DIR = os.path.dirname(SCRIPT_DIRNAME)
APP_DIR = os.path.join(PROJECT_ROOT_DIR)

sys.path.insert(0, APP_DIR)

from app import AppName
```

Of course, we need to test, but to do so, we also need to define our context.

### python_unittest.py
```python
# -*- coding: utf-8 -*-
"""
This module tests functions
"""
import unittest

from context import *


class MyTestClass(unittest.TestCase):
    """
    A simple class for testing
    """

    def setUp(self):
        self.do_tests = True
        self.app = AppName()

    def test_success(self):
        """Test this function!"""
        do_tests = self.do_tests
        self.assertEqual(do_tests, True)


if __name__ == '__main__':
    unittest.main()
```

Here, we define our unittest class.

### python.dockerfile
```Dockerfile
# base image
FROM python:3.8.0-slim as base_image
LABEL stage=appname_base_image
RUN apt-get update \
  && apt-get install gcc -y \
  && apt-get clean

WORKDIR app
COPY ./README.md /app/README.md
COPY ./requirements.txt /app/requirements.txt
COPY ./app /app
COPY ./setup.py /app/setup.py

RUN pip install --upgrade pip
RUN pip install --user --editable .

# production image
FROM python:3.8.0-slim as app_pod
COPY --from=base_image /root/.local /root/.local
COPY --from=base_image /app /app
WORKDIR app
ENV PATH=/root/.local/bin:$PATH
ENTRYPOINT ["app"]
```

For containerizing our project, we define a simple Dockerfile.

## Step 2.

Create a new bash function in a file, `new-py-project.sh`.

Copy and paste this code:

```bash
#!/bin/sh

function startvenv(){
  if [ ! -f ./pyvenv.cfg ]; then
    python -m venv ./
    mkdir ./app
    mkdir ./app/tests
    cp $HOME/Templates/python_main.py ./simple-main.py
    cp $HOME/Templates/python_cli.py ./main.py
    cp $HOME/Templates/python_class.py ./app/app.py
    cp $HOME/Templates/python_unittest.py ./app/tests/test.py
    cp $HOME/Templates/python_test_context.py ./app/tests/context.py
    cp $HOME/Templates/python.dockerfile ./Dockerfile
    cp $HOME/Templates/python_setup.py ./setup.py
    cp $HOME/Templates/python_makefile ./Makefile
    touch ./requirements.txt
    touch README.md
    echo Click >> ./requirements.txt
    echo python-dotenv >> ./requirements.txt
    source ./bin/activate
    pip install --upgrade pip
    pip install -r requirements.txt
    pip install --editable .
    app
    deactivate
  fi
}
```

## Step 3.

In your `.bash_aliases` file, put the following line,

```bash
alias newvenv='source $HOME/Utilities/new-python-project.sh && startvenv'
alias entervenv='source bin/activate'
```

Now, source your .bash_aliases file, and test out your automation pipeline:

```bash
mkdir new-py-project-test && cd new-py-project-test && \
  newvenv && \
  entervenv  
```

You should be able to run the ```app``` command from without your new directory.

Now, don't waste any more time. Start developing!
