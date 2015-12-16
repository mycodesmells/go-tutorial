# Go Lang - Getting Started

Recently I have started my journey with Go language. It seemed tempting for a long time, and recently I got the opportunity to try it out in a commercial project. Like any language of tool, the first steps are always difficult. Fortunately I have been working with more experience people, so that I got comfortable with it pretty fast. This post will help you get your feet wet and allow you to create your very first, and hopefully useful, app.

### Download

To get Go-ing (pun intended), you obviously must have Go installed on your machine. It's quite easy, as long as you follow the instructions on the [official page](https://golang.org/doc/install):

	wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
	tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz

At this point you have Go installed on your machine. This is also the first place I ran into trouble. I started by installing Go via apt-get on my Linux Mint, but it turned out to be version 1.2.1. When I realized that, I decided to install it manually. Everything was perfect until I tried to run of compile my first script, because a version installed via apt-get appeared first on my `$PATH`, and it kept failing my builds (due to some non-compatible changes between versions). So after that I now remember to set by Go executable before existng `$PATH`:

	export PATH=/usr/local/go/bin:$PATH

Also, you might consider naming the path to Go dir a `$GOROOT`, as it is often mentioned in articles or discussions on the language.

### Configuration

At this point you are able to run your scrips with `go run script.go`, but it's not good enough for your development. Within your scripts you probably are going to use some external libraries. To make it possible, you must another environmental variable - `$GOPATH`. This is used as a root directory for both your code, and the code for your dependencies. Most common place to store your code is `$HOME/go/src`, so you should set your `$GOPATH` accordingly:

	export GOPATH=$HOME/go

Note that there is no `src` at the end, as it is added during dependency search by Go itself.

### Creating executables

You are good to go (pun intended again) as long as you are just playing around with Go. But the moment you want to have your program running as a standalone executable, you need the third variabe - `$GOBIN`. It represents the directory, in which the executables will be created. Again, we'll create it in the most popular location:

	$GOBIN=$HOME/go/bin

or
	$GOBIN=$GOPATH/bin

You should consider adding your `$GOBIN` to your `$PATH`. That way you can immediately use your scripts from anywhere in your system.

### Summary

It's obviously best to have your Go configuration set up every time you enter your console. To sum up, this is what you should have configured in your console startup file (`~/.bashrc` in my case):

	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export GOBIN=$GOPATH/bin
	export PATH=$GOROOT/bin:$PATH:$GOBIN
	