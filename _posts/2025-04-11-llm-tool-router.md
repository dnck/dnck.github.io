---
layout: post
title: Route User Prompts to Subagents
author: dnck
date: 2025-04-11
comments: true
analytics: true
description: How do I route a user’s prompt to a subset of my LLM agents based on the tools that are available to them?
category: blog
show_excerpt: true
excerpt: How do I route a user’s prompt to a subset of my LLM agents based on the tools that are available to them?
---
**Problem Statement** 

Your open-weight LLM has an easy time selecting a tool given a user’s prompt when the number of tools is low, but when you increase the number of available tools by some factor, your LLM starts making rookie mistakes. 

**Solutions**

Here’s is a solution you can try:

**Solution #1**

Before prompting your model, you’re gonna have to group your tools into functionally distinct categories. These categories will form your tool catalogue table of contents. After you have your tool catalogue, craft a system prompt that instructs the model to select one or more of your tool categories from your catalogue’s TOC. This system prompt can be called your ToolRouter or something similar. Ok, now it’s time to prompt your ToolRouter.

Given a user’s query, you can do a binomial experiment using your ToolRouter. Query ToolRouter _n times_ with varying model parameters. Aggregate the results, and select the most commonly chosen tool categories for the next step: reveal the tools from the chosen categories to your model.

For each tool category, start a new chat session, and send the user’s prompt to your model with the category’s available tools. Parse the output as usual looking for valid tool calls, call the tools and provide the output back to the model in the same session. Finally, aggregate the results from each session’s output, and query your model once more in a new session using the user’s prompt and a system prompt augmented by the tool call results in a language suitable for instructing the model to respond to the original user’s query.

**Mo’ Problems**

**New Problem**: You’ve created your tool catalogue, but the number of categories has increased, and your ToolRouter has the same problem that you had before creating a tool catalogue. 

Here’s a potential solution: 

Maybe this is happening because the tool categories are not semantically distinct enough. Try running a clustering algorithm against your categories and their descriptions and rewriting them until the clusters are maximally distinct. Iterate, iterate, iterate. Once you’re satisfied with the distinctness of your tool categories, do the same for the tools within your category. The end result should be a tool catalogue with maximally distinct TOC and within each category, maximally distinct tools. In a follow up step, you can run RAG against your tool catalogue and the user’s prompt to supplement the earlier solution using the ToolRouter.











