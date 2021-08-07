#!/bin/bash

POST_TITLE="$(date +'%Y-%m-%d-')""${1}"".md"
NEW_POST_FILE="./_posts/""$POST_TITLE"
touch "$NEW_POST_FILE"

cat << EOF >> "$NEW_POST_FILE"
---
layout: post
title: $1
author: dc
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
