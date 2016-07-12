# cdgo

![Build Status](http://ci.engineerbetter.com/api/v1/pipelines/cdgo-cli/jobs/test/badge)

`cd`s to nested directories in either `$GOPATH` and `~/workspace/`, which may be handy if you regularly work on both Go projects and projects written in other languages.

## Usage

* `cdgo some-project` to `cd` into a directory somewhere in `$GOPATH/src`
* `cdwork some-other-project` to `cd` into a directory somewhere in `~/workspace`

`cdgo` ignores any directories called `vendor`. If you'd like this to be optional behaviour, please raise a GitHub issue.

## Installation

* `go get github.com/EngineerBetter/cdgo-cli` or [download a release](https://github.com/EngineerBetter/cdgo-cli/releases/latest) and save it in your `PATH` as `cdgo-cli`
* `cdgo-cli -install=$HOME/.bashrc` (see below)
* `source $HOME/.bashrc` to start pick up the new functions

### What does the install argument do?

A child process can't change its parent's working directory, so this command appends two Bash functions to the given file. You can provide a different path to `-install` if you'd rather not use `.bashrc`.

## Waffle

Directories are listed lexicographically, and each level of the directory tree is searched before descending. This means higher-level results are favoured, reducing the likelihood of `cd`ing into a submodule or vendored dependency.

Inspired by [Pivotal's Bash functions](https://github.com/cloudfoundry-incubator/garden-linux/wiki/Garden-development-workstation-setup), and hopefully a bit faster.