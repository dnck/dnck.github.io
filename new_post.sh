#!/bin/bash

export POST_TITLE=$(date +'%Y-%m-%d-')$1'.md'

touch ./_posts/$POST_TITLE

cat << EOF >> "./_posts/$POST_TITLE"
---
layout: post
title: $1
author: ea
date: $(date +'%Y-%m-%d')
comments: true
analytics: true
keywords:
description:
category: blog
show_excerpt: true
excerpt:
---
EOF
