# cdgo

[![Build Status](https://travis-ci.org/EngineerBetter/goto.svg?branch=master)](https://travis-ci.org/EngineerBetter/goto)

`cd`s to nested directories in either `$GOPATH` and `~/workspace/`, in order to make it easier to find things.

## To use

* `go get github.com/EngineerBetter/cdgo/goto`
* `go get github.com/EngineerBetter/cdgo/workto`
* Wang this in your `.bashrc`:

```
# https://github.com/EngineerBetter/cdgo
function cdgo {
  cd `goto "$@"`
}
function cdwork {
  cd `workto "$@"`
}
```

* `exec bash -l` to start a new session
* `cdgo some-project`
* `cdwork some-other-project`

Directories are listed lexicographically, and each level of the directory tree is searched before descending. This means higher-level results are favoured, reducing the likelihood of `cd`ing into a submodule or vendored dependency.

Inspired by [Pivotal's Bash functions](https://github.com/cloudfoundry-incubator/garden-linux/wiki/Garden-development-workstation-setup), but hopefully a bit faster.