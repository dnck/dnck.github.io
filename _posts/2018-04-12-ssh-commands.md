---
layout: post
title: SSH commands
date: 2018-04-12
author: dc
comments: true
analytics: false
keywords: ssh, commands
description: simple ssh commands, no bullshit
tag: ssh, commands, KISS
category: ssh, RPCs
---

SSH
You can use SSH to log into a different machines from your local machine. After you log in, you can access files, run code interpreters, etc. All of the commands below assume you're using a mac/unix like terminal with ssh in executable from some /bin folder.

1) Login via SSH <br>
A) Connect to the root directory of a server machine:<br>
<br>

<code>
<pre>
$ ssh username@server-address
$ password: ************
</code>
</pre>

B) Alternative login method:***

<code>
<pre>
$ ssh username@ip-address
$ password: ************
</code>
</pre>

You can log out by typing,

<code>
<pre>
crtl + D
</code>
</pre>

2) Upload files and directories

A) Transfer a single file from local machine to remote server:

<code>
<pre>
$ scp absolute/path/to/file.fileEnding username@server-address:/path/on/the/server/file.fileEnding
</code>
</pre>

B) Transfer an entire directory from local machine to server:

<code>
<pre>
$ scp -r absolute/path/to/the/directory usernameu@server-address:/path/on/the/remote/server/file.fileEnding]
</code>
</pre>

3) Download and directories

A) Download a single file from some directory of the remote server:

<code>
<pre>
$ scp username@server-address:/file.fileending /absolute/path/on/local/machine/file.fileending
```
B) Download a single file from a subdirectory of the remote server:
<code>
<pre>
$ scp username@server-address:/subdirectory/sub-subdirectory/file.fileending /abs/path/on/local/machine/file.fileending
</code>
</pre>

C) Download a single subdirectory and its contents from the server:
<code>
<pre>
$ scp -r username@server-address:/home/ /abs/local/path/for/new/directory/on/local/machine
</code>
</pre>

D) Download everything within the root directory:
<code>
<pre>
$ scp -r username@server-address:/ /abs/local/path/for/new/directory/on/local/machine
</code>
</pre>

NOTES
Permissions for downloading files and directories can be accessed in SSH or FileZilla and changed.

The userID on the server for SSH access will belong to a group. It's probably a good idea to check your group assignment.

View files and directory permissions on SSH with:

<code>
<pre>
$ ls -la
</code>
</pre>

Most shared hosting services suck.

The root directory of a username on a shared hosting server will end up as root for the username. Anything lower than that in the hierarchy will usually be unreadable. Just check on the group assigned as owner, and you will find out whether you can change it with the CHMOD commands above. What this means in practice is that access to any of the server side scripting languages, e.g. python, ruby, etc., pretty much anything other than php, will be impossible without paying the shitty host server (READ: Godaddy).
