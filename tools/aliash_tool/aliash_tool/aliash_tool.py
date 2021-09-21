# -*- coding: utf-8 -*-
"""app Description

This module does great things.
"""
import os
import tempfile
import subprocess
import shutil

# Implementation constants

BASH_SCRIPT_HEADER = """#!/bin/bash
#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
#     FILE:{}
#     DESCRIPTION:{}
#     USAGE:{}
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
        )
        f.write("{}".format(header))
    os.chmod(filename, 0o755)

# Classes, methods, functions, and variables
class AliashTool():
    """AliashTool is the default class for app.

    As a default, the __init__ method does not set up any class attributes.
    However, if it does, you should follow the PEP-8 conventions and document
    them as shown below.

    Args:
        msg (str): Human readable str describing the exception.
        code (:obj:`int`, optional): Error code.

    Attributes:
        msg (str): Human readable str describing the exception.
        code (int): Exception error code.

    """
    def __init__(self, home_dir, script_dir, bash_aliases_file):
        self.home_dir = home_dir
        self.script_dir = script_dir
        self.alias_definition_file = bash_aliases_file

    def join_home_dir(self, filename) -> str:
        return os.path.join(self.home_dir, filename)

    def join_script_dir(self, filename) -> str:
        return os.path.join(self.script_dir, filename)

    def _get_current_scripts_in_script_dir(self) -> list:
        return [os.path.join(self.script_dir, i) for i in os.listdir(
            self.script_dir) if i.endswith(".sh")]

    def _get_current_alias_definitions_from_file(self) -> list:
        alias_defintions = [line.strip() for line in read_file(
            self.alias_definition_file).split("\n") if not line == ""
        ]
        return alias_defintions

    def _get_current_aliases_in_alias_definition_file(self) -> list:
        alias_definition_file_aliases = []
        alias_defs = self._get_current_alias_definitions_from_file()
        for a in alias_defs:
            if "=" in a:
                alias_def = a.split("=")
                alias = alias_def[0].split("alias ")[1]
                alias_definition_file_aliases.append(alias)
        return alias_definition_file_aliases

    def _get_current_scripts_in_alias_definition_file(self) -> list:
        alias_definition_file_scripts = []
        alias_defs = self._get_current_alias_definitions_from_file()
        for a in alias_defs:
            if "=" in a:
                alias_def = a.split("=")
                filename = alias_def[1]
                alias_definition_file_scripts.append(filename)
        return alias_definition_file_scripts

    def _get_db(self) -> dict:
        """Returns {'alias': 'alias_script_path'}"""
        db = {}
        alias_defs = self._get_current_alias_definitions_from_file()
        for a in alias_defs:
            if "=" in a:
                alias_filename = a.split("=")
                alias = alias_filename[0][6:].split(" ")[0]
                filename = alias_filename[1]
                if db.get(alias) is None:
                    db[alias] = filename
                else:
                    print("ERROR: duplicate alias")
        return db

    def _format_alias_definition(self, alias) -> str:
        return "alias {}={}".format(
            alias,
            self.join_script_dir(alias+".sh")
        )

    def remove_file(self, filename):
        old_script_dir = os.path.join(self.home_dir, "Utilities/tmp")
        new_filename = filename.replace(self.script_dir, old_script_dir)
        shutil.move(filename, new_filename)

    def _clean_script_dir(self):
        scripts_in_dir = self._get_current_scripts_in_script_dir()
        scripts_in_alias_file = \
            self._get_current_scripts_in_alias_definition_file()
        for s in scripts_in_dir:
            if not s in scripts_in_alias_file:
                # actually just renames it
                self.remove_file(s)

    def _remove_alias_definition(self, alias):
        """Delete an alias definition from .bash_aliases"""
        remove_alias = self._format_alias_definition(alias)
        alias_definitions = self._get_current_alias_definitions_from_file()
        alias_definitions.remove(remove_alias)
        f = tempfile.NamedTemporaryFile(mode='w+t', delete=False)
        new_file = f.name
        for line in alias_definitions:
            if not line == "":
                f.write(line+"\n")
        f.close()
        shutil.move(
            self.alias_definition_file,
            self.alias_definition_file+".bak"
        )
        shutil.move(new_file, self.alias_definition_file)
        self._clean_script_dir()

    def _append_bash_alias_file(self, new_alias):
        """Append an alias definition from .bash_aliases"""

        new_alias_definition = self._format_alias_definition(new_alias)
        alias_definitions = self._get_current_alias_definitions_from_file()

        if not new_alias_definition in alias_definitions:
            alias_definitions.append(new_alias_definition)

        alias_definitions.sort()

        f = tempfile.NamedTemporaryFile(mode='w+t', delete=False)
        new_file = f.name
        for line in alias_definitions:
            if not line == "":
                f.write(line+"\n")
        f.close()

        shutil.move(self.alias_definition_file,
            self.alias_definition_file+".bak"
        )
        shutil.move(new_file, self.alias_definition_file)

    def _is_alias_in_script_dir(self, alias) -> bool:
        return os.path.isfile(self.join_script_dir(alias+".sh"))

    def _is_alias_in_alias_definition_file(self, alias) -> bool:
        current_aliases = self._get_current_aliases_in_alias_definition_file()
        if alias not in current_aliases:
            return False
        return True

    def add_alias(self, alias):
        """Create a new alias .sh file in the script_dir and add it to
        .bash_aliases file
        Returns:
            True
        """
        if self._is_alias_in_script_dir(alias):
            print("ERROR: alias already exists with that name")
            return True
        else:
            self._append_bash_alias_file(alias)
            new_filename = self.join_script_dir(alias+".sh")
            create_bash_script(new_filename)
            print("SUCCESS: added new alias file to script dir")
        return True

    def remove_alias(self, alias):
        """Remove an existing alias definition and its script file
        Returns:
            True
        """
        # only add the alias if it does not exist as alias and there's
        # not a filename already in the script dir
        db = self._get_db()
        old_alias_file = db.get(alias)
        if not old_alias_file is None:
            self._remove_alias_definition(alias)
            print("SUCCESS: removed old alias from .bash_aliases")
        else:
            print("ERROR: alias does not exist with that name")
        return True

    def help_alias(self, alias):
        """Show the help str from an alias definition

        Returns:
            True

        """
        # only show help if the alias does exist
        db = self._get_db()
        if db.get(alias) is None:
            print("ERROR: alias key {} not in db".format(alias))
        else:
            try:
                script = read_file(db.get(alias))
                print(script)
            except:
                print("ERROR: reading filename {}".format(db.get(alias)))
        return True

    def find_alias(self, tag) -> dict:
        """Find alias with a tag
        Returns:
            True
        """
        db = self._get_db()
        found_aliases = {}
        for alias in db:
            if tag in alias:
                found_aliases.update({alias: db[alias]})
        return found_aliases

    def edit_alias(self, alias):
        """Edit an alias

        Returns:
            True

        """
        if not self._is_alias_in_script_dir(alias):
            print("ERROR: alias script does not exist with that name")
            return True
        # only show edit if the alias does exist
        db = self._get_db()
        if db.get(alias) is None:
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
                print("ERROR: writing filename {}".format(alias+".sh"))
        return True

    def rename_alias(self, old_name, new_name):
        """Rename an alias

        Returns:
            True

        """
        if self._is_alias_in_script_dir(new_name):
            print("ERROR: alias already exists with that name in script dir")
            return True
        if not self._is_alias_in_script_dir(old_name):
            print("ERROR: alias does not exist with that name in script dir")
            return True
        if not self._is_alias_in_alias_definition_file(old_name):
            print("ERROR: alias does not exist with that name in alias file")
            return True
        self.add_alias(new_name)
        shutil.copy(
            self.join_script_dir(old_name+".sh"),
            self.join_script_dir(new_name+".sh")
        )
        self.remove_alias(old_name)
        return True

    def test_aliash_tool(self):
        """Class methods are similar to regular functions.

        Returns:
            True

        """
        return True
