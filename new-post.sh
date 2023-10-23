#!/bin/bash

# Initialize variables with default values
keywords=""
description=""
excerpt=""
POST_TITLE=""

# Function to display usage information
usage() {
    echo "Usage: $0 <POST_TITLE> [-k <arg>] [-d <arg>] [-e <arg>]"
    echo "  <POST_TITLE>: The title of the post"
    echo "  -k: Optional keywords for the post"
    echo "  -d: Optional description for the post"
    echo "  -e: Optional excerpt to show for the post"
    exit 1
}

# Check if there's a required argument
if [ $# -eq 0 ]; then
    echo "Error: Required argument is missing."
    usage
fi

# Assign the required argument
POST_TITLE="$1"

# Shift the parsed required argument out of the argument list
shift

# Parse command-line options
while getopts "k:d:e:" opt; do
    case $opt in
        k)
            keywords="$OPTARG"
            ;;
        d)
            description="$OPTARG"
            ;;
        e)
            excerpt="$OPTARG"
            ;;
        \?)
            echo "Invalid option: -$OPTARG" >&2
            usage
            ;;
    esac
done

# # Your script logic goes here
# echo "Required Argument: $(date +'%Y-%m-%d-')$POST_TITLE.md"
# echo "Optional Flag 1: $keywords"
# echo "Optional Flag 2: $description"
# echo "Optional Flag 3: $excerpt"

POST_TITLE="$(date +'%Y-%m-%d-')""${POST_TITLE}"".md"
NEW_POST_FILE="./_posts/""$POST_TITLE"

# description
function create_post() {
local keywords="$1"
local description="$2"
local excerpt="$3"
touch "$NEW_POST_FILE"
cat << EOF >> "$NEW_POST_FILE"
---
layout: post
title: $POST_TITLE
author: dnck
date: $(date +'%Y-%m-%d')
comments: true
analytics: true
keywords: $keywords
description: $description
category: blog
show_excerpt: true
excerpt: $excerpt
---
EOF
}

function main() {
  local keywords="$1"
  local description="$2"
  local excerpt="$3"

  if [ -n "$keywords" ] && [ -n "$description" ] && [ -n "$excerpt" ]; then
    create_post "$keywords" "$description" "$excerpt"

  elif [ -n "$keywords" ] && [ -n "$description" ]; then
    create_post "$keywords" "$description" ""

  elif [ -n "$keywords" ] && [ -n "$excerpt" ]; then
    create_post "$keywords" "" "$excerpt"

  elif [ -n "$description" ] && [ -n "$excerpt" ]; then
    create_post "" "$description" "$excerpt"

  elif [ -n "$keywords" ]; then
    create_post "$keywords" "" ""

  elif [ -n "$description" ]; then
    create_post "" "$description" ""

  elif [ -n "$excerpt" ]; then
    create_post "" "" "$excerpt"

  else
    create_post "" "" ""
  fi
}

main "$keywords" "$description" "$excerpt"
