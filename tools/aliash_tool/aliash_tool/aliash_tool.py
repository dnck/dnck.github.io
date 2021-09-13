# -*- coding: utf-8 -*-
"""app Description

This module does great things.
"""
import os
import tempfile
import subprocess
import shutil

# Implementation constants
SCRIPT_DIRNAME, SCRIPT_FILENAME = os.path.split(os.path.abspath(__file__))
PROJECT_ROOT_DIR = os.path.dirname(SCRIPT_DIRNAME)

BASH_SCRIPT_HEADER = """#!/bin/bash
#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
#     FILE:{}
#     DESCRIPTION:{}
#     USAGE:{}
#     AUTHOR:{}
#     DATE:{}
#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
"""

def read_file(filename):
    with open(filename, "r") as f:
        s = f.read()
    return s

def create_bash_script(filename):
    with open(filename, "w") as f:
        header = BASH_SCRIPT_HEADER.format(
            filename,
            "",#DESCRIPTION
            "",#USAGE
            "",#AUTHOR
            "",#DATE
        )
        f.write("{}".format(header))

def update_file(filename):
    with open(filename, "r") as f:
        s = f.read()
    print(s)

def remove_file(filename):
    with open(filename, "r") as f:
        s = f.read()
    print(s)

def rename_file(filename):
    with open(filename, "r") as f:
        s = f.read()
    print(s)

# Classes, methods, functions, and variables
class AliashTool():
    """AliashTool is the default class for app.

    As a default, the __init__ method does not set up any class attributes.
    However, if it does, you should follow the PEP-8 conventions and document
    them as shown below.

    Args:
        msg (str): Human readable string describing the exception.
        code (:obj:`int`, optional): Error code.

    Attributes:
        msg (str): Human readable string describing the exception.
        code (int): Exception error code.

    """
    def __init__(self, home_dir, script_dir):
        self.home_dir = home_dir
        self.script_dir = script_dir
        self.bash_aliases_file = self.join_home_dir(".bash_aliases")
        self._get_current_aliases()
        pass

    def join_home_dir(self, filename):
        return os.path.join(self.home_dir, filename)

    def join_script_dir(self, filename):
        return os.path.join(self.script_dir, filename)

    def test_aliash_tool(self):
        """Class methods are similar to regular functions.

        Returns:
            True

        """
        return True

    def _remove_bash_aliases(self, alias):
        remove_alias = "alias {}={}".format(alias, self.join_script_dir(alias+".sh"))
        bash_aliases = [i.strip() for i in read_file(
            self.bash_aliases_file).split("\n") if not i == ""
        ]
        bash_aliases.remove(remove_alias)
        f = tempfile.NamedTemporaryFile(mode='w+t', delete=False)
        new_file = f.name
        for line in bash_aliases:
            if not line == "":
                f.write(line+"\n")
        f.close()
        shutil.move(self.bash_aliases_file, self.bash_aliases_file+".bak")
        shutil.move(new_file, self.bash_aliases_file)

    def _append_bash_aliases(self, new_alias):
        new_line = "alias {}={}".format(new_alias, self.join_script_dir(new_alias+".sh"))
        bash_aliases = [i.strip() for i in read_file(
            self.bash_aliases_file).split("\n") if not i == ""
        ]
        if not new_line in bash_aliases:
            bash_aliases.append(new_line)
        bash_aliases.sort()
        f = tempfile.NamedTemporaryFile(mode='w+t', delete=False)
        new_file = f.name
        for line in bash_aliases:
            if not line == "":
                f.write(line+"\n")
        f.close()
        shutil.move(self.bash_aliases_file, self.bash_aliases_file+".bak")
        shutil.move(new_file, self.bash_aliases_file)

    def _get_current_aliases(self) -> dict:
        bash_aliases = read_file(self.bash_aliases_file)
        db = {}
        for line in bash_aliases.split("\n"):
            if "=" in line:
                alias_filename = line.split("=")
                alias = alias_filename[0][6:].split(" ")[0]
                filename = alias_filename[1]
                if db.get(alias) is None:
                    db[alias] = filename
                else:
                    print("ERROR: duplicate alias")
        self.db = db

    def add_alias(self, filename):
        """Create a new alias

        Returns:
            True

        """
        if not filename.endswith(".sh"):
            print("ERROR: new alias filename must be a bash script")
            return
        # only add the alias if it does not exist as alias and there's
        # not a filename already in the script dir
        new_alias = self.db.get(filename[:-3])
        if new_alias is None:
            if not filename in os.listdir(self.script_dir):
                new_alias = filename[:-3]
                new_filename = self.join_script_dir(filename)
                create_bash_script(new_filename)
                os.chmod(new_filename, 0o755)
                self.db[new_alias] = new_filename
                self._append_bash_aliases(new_alias)
                self._get_current_aliases()
                print("SUCCESS: added new alias file to script dir")
            else:
                print("ERROR: filename already exists in script_dir")
        else:
            print("ERROR: alias already exists with that name")
        return True

    # TODO (dnck) implement
    def remove_alias(self, alias):
        """Remove an existing alias

        Returns:
            True

        """
        # only add the alias if it does not exist as alias and there's
        # not a filename already in the script dir
        old_alias_file = self.db.get(alias)
        if not old_alias_file is None:
            self._remove_bash_aliases(alias)
            self._get_current_aliases()
            print("SUCCESS: removed old alias from .bash_aliases")
        else:
            print("ERROR: alias does not exist with that name")
        return True

    def help_alias(self, alias):
        """Show the help string from an alias definition

        Returns:
            True

        """
        # only show help if the alias does exist
        if self.db.get(alias) is None:
            print("ERROR: alias key {} not in db".format(alias))
        else:
            try:
                script = read_file(self.join_script_dir(alias+".sh"))
                print(script)
            except:
                print("ERROR: reading filename {}".format(alias+".sh"))
        return True

    def find_alias(self, tag):
        """Find alias with a tag

        Returns:
            True

        """
        found_aliases = {}
        for alias in self.db:
            if tag in alias:
                found_aliases.update({alias: self.db[alias]})
        if len(found_aliases):
            for k,v in found_aliases.items():
                print(k, v)
        else:
            print("no aliases found")
        return True

    # TODO (dnck) implement
    def tag_alias(self, alias, tag):
        """Tag an alias

        Returns:
            True

        """
        # only tag alias if the alias does exist
        if self.db.get(alias) is None:
            print("no")
        else:
            print("ok")
        return True

    def edit_alias(self, alias):
        """Edit an alias

        Returns:
            True

        """
        # only show edit if the alias does exist
        if self.db.get(alias) is None:
            print("ERROR: alias key {} not in db".format(alias))
        else:
            try:
                filename = self.join_script_dir(alias+".sh")
                script = read_file(filename)
                f = tempfile.NamedTemporaryFile(mode='w+t', delete=False)
                n = f.name
                f.write(script)
                f.close()
                subprocess.call(['nano', n])
                shutil.move(f.name, filename)
                os.chmod(filename, 0o755)
            except:
                print("ERROR: reading filename {}".format(alias+".sh"))
        return True

    # TODO (dnck) implement
    def rename_alias(self, old_name, new_new):
        """Rename an alias

        Returns:
            True

        """
        # only rename the alias if new name does not exist
        if self.db.get(filename[:-3]) is None:
            print("ok")
        else:
            print("no")
        return True