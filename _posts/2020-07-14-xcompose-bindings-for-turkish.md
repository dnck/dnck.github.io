---
layout: post
title: XCompose bindings for Turkish letters
author: dc
date: 2020-07-14
comments: true
analytics: true
keywords: ubuntu, xcompose, turkish
description:
category: blog
show_excerpt: true
excerpt: XCompose key bindings for Turkish letters on an English QWERTY keyboard.
---

It's easy to get quick access to letters in foreign alphabets on a
English QWERTY keyboard on Ubuntu. Just follow the instructions below (written
for Turkish characters) and modify as needed.

**Step 1.**
Install tweaks:
```
sudo apt install gnome-tweak-tool
```

**Step 2**
Open gnome-tweaks and select a compose key. I use the right-ctrl key.

**Step 3**
Copy & paste the text below into `~/.XCompose`.

```
# This file defines custom Compose sequences for Unicode characters

# Import default rules from the system Compose file:
include "/usr/share/X11/locale/en_US.UTF-8/Compose"

# To put some stuff onto compose key strokes:
<Multi_key> <C> <C> <C> : "Ç" 00c7
<Multi_key> <c> <c> <c> : "ç" 00e7
<Multi_key> <G> <G> <G> : "Ğ" 011e
<Multi_key> <g> <g> <g> : "ğ" 011f
<Multi_key> <I> <I> <I> : "İ" 0130
<Multi_key> <i> <i> <i> : "ı" 0131
<Multi_key> <O> <O> <O> : "Ö" 00d6
<Multi_key> <o> <o> <o> : "ö" 00f6
<Multi_key> <S> <S> <S> : "Ş" 015e
<Multi_key> <s> <s> <s> : "ş" 015f
<Multi_key> <U> <U> <U> : "Ü" 00dc
<Multi_key> <u> <u> <u> : "ü" 00fc
```

That's it. Log-out and log back in and your new bindings should take effect.
To test, enter your compose key, and then enter three "g" keys. You should see
the yumuşak ge (ğ).
