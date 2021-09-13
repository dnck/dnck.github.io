# -*- coding: utf-8 -*-
"""aliash_tool Description

This is the command line interface for aliash_tool. The structure of this program
follows the structure of Click command line interface. You can read more about
Click here, https://click.palletsprojects.com/en/7.x/
"""
import click

from aliash_tool import aliash_tool


class CommandLineApp():
    """CommandLineApp is the default command line client for aliash_tool.

    It uses the Click module instead of argparse.
    """
    def __init__(self):
        self.cli = aliash_tool.AliashTool("/home/dnck", "/home/dnck/Utilities")

@click.group()
@click.pass_context
def cli(ctx):
    """
    aliash_tool manages your .bash_aliases!
    """
    ctx.obj = CommandLineApp()


@cli.command()
@click.pass_context
def test(ctx):
    """
    Test all methods of aliash_tool
    """
    assert ctx.obj.cli.test_aliash_tool()

@cli.command()
@click.argument('filename')
@click.pass_context
def add_alias(ctx, filename):
    """
    Create a new [FILENAME] and put its alias in .bash_aliases
    """
    assert ctx.obj.cli.add_alias(filename)

@cli.command()
@click.pass_context
@click.argument('alias')
def remove_alias(ctx, alias):
    """
    Remove an [ALIAS] from .bash_aliases
    """
    assert ctx.obj.cli.remove_alias(alias)

@cli.command()
@click.pass_context
@click.argument('alias')
def edit_alias(ctx, alias):
    """
    Edit an [ALIAS] in the script_dir (requires nano)
    """
    assert ctx.obj.cli.edit_alias(alias)

@cli.command()
@click.pass_context
@click.argument('alias')
@click.argument('new_name')
def rename_alias(ctx, alias, new_name):
    """Rename an [ALIAS] in .bash_aliases"""
    assert ctx.obj.cli.rename_alias(alias, new_name)

@cli.command()
@click.pass_context
@click.argument('alias')
@click.argument('tag')
def tag_alias(ctx, alias, tag):
    """Tag an [ALIAS] in .bash_aliases"""
    assert ctx.obj.cli.tag_alias(alias, tag)

@cli.command()
@click.pass_context
@click.argument('tag')
def find_alias(ctx, tag):
    """
    Find an alias in .bash_aliases using a [TAG]
    """
    assert ctx.obj.cli.find_alias(tag)

@cli.command()
@click.pass_context
@click.argument('alias')
def help_alias(ctx, alias):
    """
    Display help for an [ALIAS] in .bash_aliases
    """
    assert ctx.obj.cli.help_alias(alias)

if __name__ == "__main__":
    cli()
