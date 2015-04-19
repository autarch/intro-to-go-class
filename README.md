# Installing Go

See the [Installing Go](http://golang.org/doc/install) page golang.org for
information on installing Go. This page has instructions for installing from
tarball on Linux, OSX, and FreeBSD. It also has links to an OSX package
installer, an MSI installer for Windows, and a few other options.

If you're on Ubuntu
[Jay R Wren's Go PPA](https://launchpad.net/~evarlast/+archive/ubuntu/golang1.4)
is a good option for installing Go 1.4 on your system. Fedora 21 and 22 both
include Go 1.4.

# Operating Systems

Go should work identically on Linux, OSX, and Windows systems.

The exercises repository uses symlinks in the answer directories to avoid
having to copy some files around. However, you don't need to run or edit any
code in the answer directories (they're for the instructor).

# Emacs Integration

If you're using Emacs as your editor, I highly recommend installing
[go-mode](http://www.emacswiki.org/emacs/GoMode) for syntax highlighting. You
can install this using [ELPA](http://www.emacswiki.org/emacs/ELPA). Add the
http://melpa.milkbox.net/packages package repo in order to get the latest
go-mode package.

I also suggest binding a key to execute `gofmt` on the current buffer. Here's
a snippet that binds M-t:

    (add-hook 'go-mode-hook
              (lambda ()
                (local-set-key "\M-t" `gofmt)))


# Vim Integration

Here are some useful Vim tools for Go:

* https://github.com/fatih/vim-go - syntax highlighting, auto gofmt on save, and more
* https://github.com/scrooloose/syntasti - syntax checking as you type
* See https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins for more

# Sublime Text Integration

Check out the [GoSublime](https://github.com/DisposaBoy/GoSublime) plugin
collection if you're using Sublime Text. It supports code completion, syntax
checking as you type, quick formatting, and much more.

# Getting Set Up for the Class

First, you'll want to check this repo out. Then run the appropriate binary for
your platform from the `setup` directory.

If you're on a platform which doesn't have a binary, all it does is the
following:

* Makes a `$GOPATH` root directory under `$HOME/go`
* Creates `bin`, `pkg`, and `src` directories under that `$GOPATH`
* Creates `src/github.com/autarch` and clones `https://github.com/autarch/intro-to-go-class-exercises.git` from that subdirectory
* Creates `src/github.com/stretchr` and clones `https://github.com/stretchr/testify.git from that subdirectory
* Tells you to set your `$GOPATH` environment variable to `$HOME/go` if it's not already set

Note that you can use a different directory for your `$GOPATH` if you prefer.
