# cdgo

[![Build Status](https://travis-ci.org/EngineerBetter/cdgo.svg?branch=master)](https://travis-ci.org/EngineerBetter/cdgo)

`cd`s to nested directories in either `$GOPATH` and `~/workspace/`, which may be handy if you regularly work on both Go projects and projects written in other languages.

## To install

* `go get github.com/EngineerBetter/cdgo/goto`
* `goto -install=$HOME/.bashrc` adds two Bash functions to the `.bashrc`
* `exec bash -l` to start a new session and pick up the new functions

You can provide a different path to `goto -install` if you'd rather not use `.bashrc`.

## To use

* `cdgo some-project` to `cd` into a directory somewhere in `$GOPATH/src`
* `cdwork some-other-project` to `cd` into a directory somewhere in `~/workspace`

Directories are listed lexicographically, and each level of the directory tree is searched before descending. This means higher-level results are favoured, reducing the likelihood of `cd`ing into a submodule or vendored dependency.

Inspired by [Pivotal's Bash functions](https://github.com/cloudfoundry-incubator/garden-linux/wiki/Garden-development-workstation-setup), and hopefully a bit faster.