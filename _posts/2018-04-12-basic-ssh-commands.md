---
layout: post
title: Basic SSH commands
date: 2018-04-12
author: dc
comments: true
analytics: false
keywords: ssh, commands
description: simple ssh commands, no bullshit
tag: ssh, commands, KISS
category: blog
---

You can use SSH to log into a different machines from your local machine. After you log in, you can access files, run programs, etc. All of the commands below assume you're using a mac/unix like terminal with the ssh executable on your path.

**Log in & out**<br>

**Login and connect to the root directory of a server machine:**
```
$ ssh username@server-address
```
**Alternative login method:**
```
$ ssh username@ip-address
```

**Log out**
```
crtl + D
```
<hr>
<br>
**Upload files and directories**

**Transfer a single file from local machine to remote server**
```
$ scp absolute/path/to/file.fileEnding username@server-address:/path/on/the/server/file.fileEnding
```

**Transfer an entire directory from local machine to server**
```
$ scp -r absolute/path/to/the/directory usernameu@server-address:/path/on/the/remote/server/file.fileEnding]
```
<hr>
<br>
**Download and directories**
**Download a single file from some directory of the remote server**
```
$ scp username@server-address:/file.fileending /absolute/path/on/local/machine/file.fileending
```
**Download a single file from a subdirectory of the remote server**
```
$ scp username@server-address:/subdirectory/sub-subdirectory/file.fileending /abs/path/on/local/machine/file.fileending
```

**Download a single subdirectory and its contents from the server**
```
$ scp -r username@server-address:/home/ /abs/local/path/for/new/directory/on/local/machine
```

**Download everything within the root directory**
```
$ scp -r username@server-address:/ /abs/local/path/for/new/directory/on/local/machine
```
<hr>
<br>
**Some reminders**<br>
Permissions for downloading files and directories can be accessed in SSH or FileZilla and changed.

The userID on the server for SSH access will belong to a group. It's a good idea to check your group assignment.

View files and directory permissions on SSH with:
```
$ ls -la
```

The root directory of a username on a shared hosting server will end up as root for the username. Anything lower than that in the hierarchy will usually be unreadable. Just check on the group assigned as owner, and you will find out whether you can change it with the CHMOD commands above. What this means in practice is that access to any of the server side scripting languages, e.g. python, ruby, etc., pretty much anything other than php, will be impossible without paying the shitty host server (READ: Godaddy).
