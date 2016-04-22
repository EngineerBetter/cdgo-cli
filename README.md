# cdgo

[![Build Status](https://travis-ci.org/EngineerBetter/cdgo.svg?branch=master)](https://travis-ci.org/EngineerBetter/cdgo)

`cd`s to nested directories in either `$GOPATH` and `~/workspace/`, which may be handy if you regularly work on both Go projects and projects written in other languages.

## Usage

* `cdgo some-project` to `cd` into a directory somewhere in `$GOPATH/src`
* `cdwork some-other-project` to `cd` into a directory somewhere in `~/workspace`

## Installation

* `go get github.com/EngineerBetter/cdgo/goto`
* `goto -install=$HOME/.bashrc`
** Performs `go get github.com/EngineerBetter/cdgo/workto`
** Adds two Bash functions which call the Go binaries that were just installed
* `exec bash -l` to start a new session and pick up the new functions

You can provide a different path to `goto -install` if you'd rather not use `.bashrc`.

## Waffle

Directories are listed lexicographically, and each level of the directory tree is searched before descending. This means higher-level results are favoured, reducing the likelihood of `cd`ing into a submodule or vendored dependency.

Inspired by [Pivotal's Bash functions](https://github.com/cloudfoundry-incubator/garden-linux/wiki/Garden-development-workstation-setup), and hopefully a bit faster.