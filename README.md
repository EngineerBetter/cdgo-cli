# cdgo

[![Build Status](https://travis-ci.org/EngineerBetter/goto.svg?branch=master)](https://travis-ci.org/EngineerBetter/goto)

Makes it easier to `cd` between projects in `$GOPATH` and `~/workspace/`.

## To use

1. `go get github.com/EngineerBetter/goto/goto`
1. `go get github.com/EngineerBetter/goto/workto`
1. Wang this in your `.bashrc`:

```
# https://github.com/EngineerBetter/cdgo
function cdgo {
  cd `goto "$@"`
}
function cdwork {
  cd `workto "$@"`
}
```

1. `exec bash -l` to start a new session
1. `cdgo some-project`
1. `cdwork some-other-project`


Inspired by [Pivotal's Bash functions](https://github.com/cloudfoundry-incubator/garden-linux/wiki/Garden-development-workstation-setup).