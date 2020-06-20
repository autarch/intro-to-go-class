# Installing Go

See [Installing Go on golang.org](http://golang.org/doc/install) for details
on installing Go. This page has instructions for installing from a
pre-compiled tarball on Linux, OSX, and FreeBSD. It also has links to an OSX
package installer, an MSI installer for Windows, and a few other options.

If you're on Ubuntu 20.04, you can install the `golang-1.14-go` package. For
other versions of Ubuntu, I recommend the `longsleep/golang-backports` PPA:

    sudo add-apt-repository ppa:longsleep/golang-backports
    sudo apt update
    sudo apt install golang-go

For Fedora and Red Hat, check out https://go-repo.io/ for golang 1.14.x
packages.

# Operating Systems

Go should work identically on Linux, OSX, and Windows systems.

The exercises repository uses symlinks in the answer directories to avoid
having to copy some files around. However, you don't need to run or edit any
code in the answer directories (they're for the instructor).

# Emacs Integration

If you're using Emacs as your editor, I highly recommend installing
[go-mode](https://github.com/dominikh/go-mode.el) for syntax highlighting. You
can install this using [ELPA](http://www.emacswiki.org/emacs/ELPA). Add the
http://melpa.milkbox.net/packages package repo in order to get the latest
go-mode package.

I also suggest binding a key to execute `gofmt` on the current buffer. Here's
a snippet that binds M-t:

    (add-hook 'go-mode-hook
              (lambda ()
                (local-set-key "\M-t" `gofmt)))

You can also bind this to
[`goimports`](https://godoc.org/golang.org/x/tools/cmd/goimports). Install
this with `go get golang.org/x/tools/cmd/goimports`. This wraps `gofmt` and
makes sure that the list of package imports for the file is correct.

# Vim Integration

Here are some useful Vim tools for Go:

* https://github.com/fatih/vim-go - syntax highlighting, auto gofmt on save, and more
* https://github.com/vim-syntastic/syntastic - syntax checking as you type
* See https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins for more

# Sublime Text Integration

Check out the [GoSublime](https://github.com/DisposaBoy/GoSublime) plugin
collection if you're using Sublime Text. It supports code completion, syntax
checking as you type, quick formatting, and much more.

# General Editor Recommendations

Setting up your editor to integrate with `go fmt` and `go vet` makes your Go
coding experience much better. You can have it run these checks on the
currently loaded file and/or run these checks every time you save your code.

You may also want to be able to run `golint` easily, though I don't recommend
this as a default action for every save. It will often complain about issues
that you will not want to fix, especially when you write code for this class's
exercises.

# Getting Set Up for the Class

First, you'll want to check this repo out. Then run the appropriate binary for
your platform from the `setup` directory. If you want to specify the `GOPATH`
for the binary to use, simply set it on the command line:

    GOPATH=$HOME/go class-setup-darwin-amd64

Note that on OSX you **must** set the `GOPATH` variable before running the
binary.

If you're on a platform which doesn't have a binary, all it does is the
following:

* Makes a `$GOPATH` root directory under `$HOME/go`.
* Creates `$GOPATH/src/github.com/autarch` and clones
  `https://github.com/autarch/intro-to-go-class-exercises.git` from that
  subdirectory.
* Tells you to set your `$GOPATH` environment variable to `$HOME/go` if it's
  not already set.
* Tell you to add `$GOPATH/bin` to your `$PATH`.

Note that you can use a different directory for your `$GOPATH` if you prefer.

**Make sure that you set the `GOPATH` environment variable and add
`$GOPATH/bin` to your `PATH` in your preferred shell's rc file(s) such as
`~/.bashrc`, `~/.zshrc`, etc.**

**You should also set your `GOROOT` environment variable. This is the root
directory of your go installation. On Ubuntu with the `golang-1.14-go` package
this is `/usr/lib/go-1.14`.**

# Reading the Slides on Your System

Being able to browse the slides while doing the exercises will be very
helpful, so I recommend getting this set up in advance.

First, you'll need to clone this repo. Then in the checkout run the following
two commands:

    git submodule init
    git submodule update

Next you'll need the present tool. After you have your `$GOPATH` set up,
simply run:

    go get golang.org/x/tools/cmd/present

The `present` binary will be installed as `$GOPATH/bin/present`, which you
should have added to your `PATH`.

You can run `present` from the checkout directory, which contains the
`intro-to-go.slide` file. Open up the URL that `present` gives you and you'll
see a link to open the slides.
