# -*- coding: utf-8 -*-
"""
This module tests functions
"""
import unittest

from context import PrimaryClassName


class MyTestClass(unittest.TestCase):
    """
    A simple class for testing
    """

    def setUp(self):
        """Set up aliash_tool.PrimaryClassName"""
        self.do_tests = True
        self.aliash_tool = PrimaryClassName()

    def test_success(self):
        """Test this function!"""
        self.assertEqual(self.aliash_tool.default_func(), True)


if __name__ == "__main__":
    unittest.main()
