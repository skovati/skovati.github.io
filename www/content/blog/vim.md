---
title: "Vim: The World's Greatest Text Editor"
date: 2021-06-03T20:04:39-05:00
draft: true
tags:
  - vim
  - unix
  - programming
---
*Disclaimer: This blog post is not meant to teach Vim, (that might come later) just to act as inspiration to start the learning journey
if you've been convinced by this post and want to learn Vim, here are a few fantastic tutorials that helped me get started*:

- `vimtutor`, Vim's built in tutorial, if Vim is installed on your system, just run `vimtutor` for an interactive tutorial
- [Ben Awad's Tutorial](https://www.youtube.com/watch?v=IiwGbcd8S7I) - watch this first, no BS, beginner to advanced
- [DistroTube's Vim series](https://www.youtube.com/watch?v=ER5JYFKkYDg) - moderate BS, takes it slow
- And for Vim-god inspiration, check out a few [Luke Smith videos](https://www.youtube.com/watch?v=hraHAZ1-RaM) - dude himself is BS, but rocks at Vim
- And these [Vim meetup talks](https://www.youtube.com/watch?v=wlR5gYd6um0) are a good rabbit hole

Anyways... the post...

## What is a Vim anyways?
Vim stands for vi-improved. It builds off of the ideas of a previous text editor known as vi, which was created by [Bill Joy](https://en.wikipedia.org/wiki/Bill_Joy) (yeah *that* Bill Joy that worked on Unix, BSD, TCP/IP, Java, and NFS) back in 1976. Lots of programs back then were iterated on by others, since the source code was often available for modification. Vi is no exception.

## Why use Vim?
One of the main groundbreaking ideas behind vi, and therefore Vim, is the idea of *modal editing*, where you switch between modes and keys do different things in each of those modes. This is in stark contrast to editors a lot of people are familiar with like Microsoft Word or the Google Docs editor, where every qwerty key simply inputs the corresponding character into the document. In Vim, this **only** happens when in 'insert' mode. When in the other modes, like 'normal' or 'visual' these same keys you use to input characters are instead used as keybinds for movement, editing, saving, quitting, etc. The most obvious advantage to this is that you can have significantly more keybinds than non-modal editors, since you simply have more keys to work with. Microsoft Word needs to use some hand-cramp-inducing combination of keys like `Ctr-Shift-t` for simple commands like cut or paste, when in Vim this can just be one key, right by the asdfjkl; homerow. This offers more specific keybinds that can significantly increase productivity, as well as reduce hand fatigue when typing for 8+ hours a day. 

The other groundbreaking idea behind Vim (and vi) is *command composition*. We'll dive into this in the next section, but essentially Vim has different types of commands like movement, editing, and repetition, which can be composed to increase their functionality. So you can apply any editing command to a movement command to apply it to that whole chunk of text (again, we'll see examples shortly). This means that you can rapidly gain new functionality by learning just a few of the commands Vim has to offer; say you know 5 editing commands and 6 movement commands, since you can compose these by placing the editing command before the movement command, you now know 5 x 6 = 30 commands while only memorizing 5 + 6 = 11 keybinds! Having that same amount of editing power in MS Word would require learning *all* 30 keybinds.

Finally, Vim has a rich editing philosophy. A lot of the design decisions were made due to technical limitations at the time (no mouse, terminal screens only, slow computers that benefited from chaining sequential commands), but these decisions turn out to have modern day benefits as well (no back and forth between keyboard and mouse, easier on the hands and faster, composable commands allow for verbose, flexible expression of editing desires, commands offer easy automation for repetitive tasks, integrates with Unix Philosophy, etc), since good design is still good design, even 50 years later.

## Vim is an editing language, with syntax and grammar
Okay so, what are these different commands I keep talking about and why are they so powerful? Again, this isn't a tutorial, but I'll list the basics here so I can use them in examples.

### Editing commands
- `i`   - enter insert mode
- `Esc` - return to normal mode
- `d`   - delete
- `c`   - change
- `y`   - copy
- `p`   - paste

### Movement commands
- `h, j, k, l`  - move left, down, up, right
- `w`           - move to next work
- `0`           - move to beginning of line
- `$`           - move to end of line

### Text Objects
- `w`   - word


## Who is Vim for?
