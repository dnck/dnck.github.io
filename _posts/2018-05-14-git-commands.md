---
layout: post
title: git commands
author: dc
date: 2018-05-14
comments: false
analytics: false
keywords: git, commands
description: simple git commands, no bullshit
tag: git, commands, KISS
category: git
---



Here are some of the common git commands, or questions that I find myself googling a lot. Maybe they help you.


**get a git repo**
```
git clone https://github.com/HelixNetwork/helix-whitepaper.git
```

**Show local AND remote branches**
```
git branch -a
```

**checkout and/or create your local branch**
```
git checkout branch_name
```

**checkout branch from the remote NOT just your local possibly new branch**
```
git checkout origin/branch_name
```

**basic flow to add changes to the current branch your working on**
```
git add *
git commit -m "message"
git push
```
Using the * means all all your changes

**update the repo to the most recent commits**
```
git pull
```

**delete remote branch after done your work on it**
```
git push origin --delete branch_name
```

**What is the difference between upstream and downstream?**

> "You're **downstream** when you **copy** (**clone**, **checkout**, etc) from a repository. Information flowed **downstream** to you. <br>
When you make changes, you usually want to send them back **upstream** so [your changes] make it into [the] repository [...] that everyone [is] pulling from [...] This is mostly a social issue of how everyone can coordinate their work rather than a technical requirement of source control. You want to get your changes into the main project so you're not tracking divergent lines of development." <a href="https://stackoverflow.com/users/230468/dilithiummatrix">-Stackoverflow user</a>

The point here is that, if you clone a repo, and start making changes on a local branch you've set up, you're probably gonna wanna set the upstream remote as the git repo at some web address, and then send the work you're doing to that upstream.
