# -*- coding: utf-8 -*-
# pylint: skip-file
"""
This script is a testing utility used to aliash_toolend the '../aliash_tool' directory
(where the primary .py executables live) to the system path so that they can be
imported during the tests.
"""
import os
import sys


SCRIPT_DIRNAME, SCRIPT_FILENAME = os.path.split(os.path.abspath(__file__))
PROJECT_ROOT_DIR = os.path.dirname(SCRIPT_DIRNAME)
APP_DIR = os.path.join(PROJECT_ROOT_DIR)

sys.path.insert(0, APP_DIR)

from aliash_tool import PrimaryClassName
