---
layout: post
title: Lastpass chrome installer causing chrome to be "managed by organization"
author: dc
date: 2019-05-07
comments: true
analytics: true
keywords:  
description:
category: blog
---

Problem: Lastpass stopped working in Chrome and Firefox
Attempt at solution: Uninstall Firefox, download newest version, reinstall.

Outcome: Nope, there was still a problem in Firefox, which they're trying to fix.

Solution: Migrate to Chrome?

Outcome: Lastpass installer downloaded from the lastpass website on this date, but it didn't stick to chrome. The final screen of the ```install_lastpass.sh``` script launched in the browser shows an inactive "install lastpass" button.

Solution: Run the ```uninstall_lastpass.sh``` from terminal. Looks like I will have to cp my passwords? (Lastpass is on its last breathe.)

New problem: Chrome was still showing a message that it (chrome) was now managed by a system-wide policy, and I don't like allowing external websites permissions to install extensions on my browser (see ```chrome://policy```).

Problem defined: The lastpass ```uninstall_lastpass.sh``` script does not remove policies it sets for itself in directory, ```/etc/opt/chrome/policies/managed```. I supsect this is not malicious, but simply negligence out of ignorance.

Final solution: ```sudo rm /etc/opt/chrome/policies/managed/lastpass_policy.json```

Outcome: Chrome no longer shows that it is ```Managed by an organization```.
