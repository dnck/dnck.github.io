---
layout: post
title: aliash
date: 2021-09-21
---

aliash is a tool for manipulating .bash_aliases file. It is very opinionated in
that it expects all bash scripts to be stored in a `script_dir` and the alias
definitions referencing these scripts to be in a single `.bash_aliases` file.

**Install**

```
python -m . venv
source bin/activate
pip install --upgrade pip
pip install -r requirements.txt
pip install --editable .
aliash --help
```

**Usage**

```
Usage: aliash_tool [OPTIONS] COMMAND [ARGS]...

  aliash_tool manages your .bash_aliases!

Options:
  --help  Show this message and exit.

Commands:
  add     Create a new [ALIAS] and put its empty definition in .bash_aliases
  edit    Edit an [ALIAS] in the script_dir (requires nano)
  find    Find an alias in .bash_aliases using a [TAG]
  help    Display help for an [ALIAS] in .bash_aliases
  insert  Display help for an [ALIAS] in .bash_aliases
  remove  Remove an [ALIAS] from .bash_aliases
  rename  Rename an [ALIAS] in .bash_aliases
  test    Test all methods of aliash_tool
```
