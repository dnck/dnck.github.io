import json
import datetime

x = datetime.datetime.now()
today = x.strftime("%Y-%m-%d")

contributions_file = "all.json"
contribs_category = ["publications", "preprints", "oral-presentations", "poster-presentations", "press"]

with open(contributions_file, "r") as f:
    contribs = json.load(f)[0]

post = """---
layout: post
title: {}
author: ea
date: {}
comments: false
analytics: false
keywords: heart, brain, consciousness, neuroscience
description: {}
category: contributions
show_excerpt: false
---

{}
"""

for i in contribs_category:
    for x in contribs.get(i):
        title_base = x.get("title").replace(" ", "-").lower(). replace("?", "")
        title_url = "<a href='{}'>{}</a>".format(x.get("url"), x.get("title"))
        if i == "publications":
            citation = "{} ({}). {}. {}.".format(x.get("authors"), x.get("date"), title_url, x.get("journal"))
        if i == "preprints":
            citation = "{} ({}). {}. {}.".format(x.get("authors"), x.get("date"), title_url, x.get("journal"))
        if i == "oral-presentations":
            citation = "{}. {}. {}.".format(title_url, x.get("date"), x.get("conference"))
        if i == "poster-presentations":
            citation = "{}. {}. {}.".format(title_url, x.get("date"), x.get("conference"))
        if i == "press":
            citation = "{}. {}. {}.".format(title_url, x.get("source"), x.get("date"))

        new_post = post.format(title_base, today, x.get("title"), citation)
        #print(new_post)
        with open("/home/dnck/Blog/Esra-Al.github.io/_posts/{}-{}.md".format(today, title_base), "w") as f:
            f.write(new_post)
