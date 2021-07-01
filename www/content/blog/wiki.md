---
title: "A minimal Vim, Markdown, and GPG powered wiki"
date: 2021-06-20T21:40:31-05:00
draft: true
tags:
  - vim
  - unix
  - minimalism
---

## Motivation and Vimwiki
Lately, I've been keeping notes, ideas, to-do lists, and journal entries in a test based wiki, thanks to the fantastic plugin [vimwiki](https://github.com/vimwiki/vimwiki). If you're not familiar, vimwiki enables easy hyperlinking between text files marked up with either Markdown or a feature rich vimwiki syntax. This can be a collection of programming language snippets, notes from CS class, project to-do lists, a personal diary—whatever you want. You can navigate between linked files with <Enter> and <Backspace> and even convert to HTML to view in a browser. I'm actually in the process of creating my own minimal plugin with a subset of the vimwiki features I actually use, along with some other customizations, since a lot of the behavior I desire can be emulated within vanilla vim with `./relative/paths` and `gf`. Either way, a markdown/plain text based wiki offers many advantages, as it can be version controlled with git, takes up kBs of space, can be efficiently synced across computers, and encrypted for privacy. I tried other note taking apps like [Joplin](https://joplinapp.org/), [Standard Notes](https://standardnotes.org/), but why install an Electron app when we can use some Vimscript?

## Markdown
Markdown has become my go-to markup language for notes and lists. It offers the essential features of a simple markup language, while staying easy to use with its syntax. I often make use of the checkboxes and tables to keep track of lists and information, which are nice additions alongside the basics of making text **bold** or *italicized* and organizing sections and lists. Plus, since it's just plain text, unlike a Microsoft Word document, I can edit it in vim (with the help of [vim-markdown](https://github.com/plasticboy/vim-markdown)), interact with it via unix command line tools, 

An example markdown wiki
```markdown
## to-do
- [ ] wash the dishes
- [ ] take out the trash
- [X] study

## projects
### games
- [pong](./projects/pong.md)
- [space invaders](./projects/invaders.md)

### other
- [garden](./projects/garden.md)
- [blog](./projects/blog.md)

## university classes
- Algorithms
- Number Theory
- Electrodynamics
- Metaphysics
- Statistics
```

## Vim

## GPG

## POSIX shell
