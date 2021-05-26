## Bash && Extensibility
*Author: Skovati, Date: 2020-12-29*

So what *is* bash? It stands for Bourne Again Shell, and is a rewrite and play on words of the original Bourne shell (sh) written by Stephen Bourne at Bell Labs. It is mainly used as an interactive command line interpreter. When you login to a Unix-based system, or open a terminal, you'll be greeted by a prompt that looks something like `[username@hostname ~]$`. Here you can run any of the GNU system utility commands. Here is a short list of important commands you should know.

- `ls` - LiSts the files and directories in the current directory
- `cd` - Changes the Directory you are in, eg `cd downloads` would move you to your downloads directory
- `mkdir` - MaKes a DIRectory
- `man` - prefixing this to another command shows the MANual page for that command—this is *the* best way to learn about a program quickly
- `touch` - TOUCHes a file to update its metadata, or will create an empty file if it doesn't exist
- `cat` - conCATenate the files passed to it and prints them. Commonly used to view the contents of a text file quickly.
- `mv` - MoVes argument 1 to argument 2, eg `mv resume_current.pdf resume_old.pdf`
- `cp` - CoPies argument 1 to argument 2
- `rm` - ReMoves a file or directory
- `pwd` - Prints Working Directory, prints the directory you're currently in
- `nano` - GNU text editor, pass a file to it to edit, perfect for editing configuration files.

These commands will likely make up the majority of what you run, as they help you move around the filesystem and edit files on the command line. Really the best way to get familiar with these commands and more is to use the terminal as often as you can! Personally, I exclusively use the terminal to edit files—I'm currently writing this in Vim, a terminal based text editor. 

Every time you run a command in bash, it's actually calling a pre-compiled C program binary located at /bin or /usr/bin. You can see the path of the executed binary with the `which` command. For example, running `which ls` outputs /usr/bin/ls, which is the small executable that bash is calling every time I run that command. 

Bash is more than just a command line interpreter however, it's actually a full scripting language with control flow, variables, and functions. While these features *could* be used on the command line, they're most commonly used in a *shell script*, which has a .sh extension, and can be executed by the bash shell. Personally, I have dozens of small bash scripts to do things like update my system, quickly find and edit files, connect to vpns and more. A decent rule of thumb for when something should be converted to a shell script is if you find yourself running a series of commands in a row multiple times a day, turn those commands into a small shell script that bash can execute for you. For example, I'm writing this blog in Markdown, and using Pandoc to convert that to html so my web server can host it. Essentially, the output of pandoc has to be concatenated with a template.html file I have, and then I have to concatenate that index.html output with a closing article tag so it's valid html. So, instead of running that set of commands multiple times, I wrote the following shell script to speed things up. 

    #!/usr/bin/env bash

    cp -v /home/skovati/code/git/skovati.com/blog/template.html ./index.html
    pandoc $1 -f markdown -t html >> ./index.html
    echo "</article>" >> ./index.html

Each of those lines *could* just be pasted into a regular bash prompt, and it would work the same, but the script saves you a lot of time. The first line there just tells the bash interpreter to run the script. Then I copy my template html file to the current chapter folder, convert the chapters markdown to html and add it to the end of the template html file, and finally add that closing article tag. Now, I can just run `convert ch6.md` and it'll output a correctly formatted html file for my server to host. 

Bash scripting should (and probably will be) its own chapter, but I want to show off a few *really* cool things it can do. First, you should be familiar with standard input and output. Stdin is essentially a stream of text that can be sent to a program, and stdout is a stream of text that a program can output. On Unix-based systems, we can use these in combination with the Unix Philosophy of each program doing one thing and doing it well, to chain together programs. Here's an example.

Running `ls` in my current directory prints out three directories: code, documents, and downloads. Say, for example, I wanted to count how many files and directories were output by that `ls` command. Here, we can use "pipes" by placing the "|" symbol in-between `ls` and another command that counts the number of lines in a text stream, `wc -l`. So, the full bash command would look like `ls | wc -l` which would output 3, since there were just those three files. Say now I wanted to check the number of installed binary programs on my system, I could just run `ls /usr/bin | wc -l` which outputs 3490. 

This introduces the idea of extensibility. Since these GNU system utilities simply take in a text stream and then output a text stream, they can easily communicate with each other, which creates new possibilities that the developers of either program didn't have to explicitly program. Here's a more useful example of this text-based extensibility in action:

As previously mentioned, I'm writing these blog posts in Markdown, which is just a text file with some simple syntax to show **bold** words or *italics* and stuff like that. Since all the posts are text, I can run cat on all of them to output them all to a big text stream. Then, I could pipe that into another program, like word count, to see more information about all the blog posts. Running `cat ch*/*.md | wc -w` (the \*'s just mean to select every .md file) outputs 7288, so that's how many words I've written in the last 5 blog posts! Going deeper, I could count up how many times I've said the word "Linux" with a command like this `cat ch*/*.md | grep -wo Linux | wc -l` (grep is a tool for searching text streams) which outputs 90. If I was instead writing all of these in something like Microsoft Word, I'd have to open up each file manually, Ctrl-f to search for Linux, and then count up and add together all of the instances manually. *That's* the power of Linux. 

