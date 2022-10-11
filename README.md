# gograveyard

*The Go project undertaker: check go.mod dependency's health*

[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/goreapers/gograveyard/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/goreapers/gograveyard?status.svg)](https://godoc.org/github.com/goreapers/gograveyard)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/goreapers/gograveyard/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/goreapers/gograveyard/tree/master)

## Description

## Usage

To run gograveyard pass either a URL to a repo to analyze or the path to a
local repo:

```s
$ gograveyard [--debug] [--json] [url|path]
```

The JSON flag will enable the output of the report in JSON.

The debug flag will produce additional output at each step useful for debugging
issues.

## Install

Below outlines the various ways to obtain and install gograveyard.

### Via go

To download using the `go install` command run:

```shell
# Install the latest release:
$ go install github.com/goreapers/gograveyard/cmd/gograveyard@latest

# Install at tree head:
$ go install github.com/goreapers/gograveyard/cmd/gograveyard@master

# Install at a specific version or pseudo-version:
$ go install github.com/goreapers/gograveyard/cmd/gograveyard@v1.1.0
```

The executable object file location will exist at `${GOPATH}/bin/gograveyard`

### From binary

Download the [latest release][latest_release]
of gograveyard for your platform and extract the tarball:

```shell
wget gograveyard<version>_<os>_<arch>.tar.gz
tar zxvf gograveyard<version>_<os>_<arch>.tar.gz
```

The tarball will extract the readme, license, and the pre-compiled binary.

[latest_release]: https://github.com/goreapers/gograveyard/releases/latest

### From source

To build and install directly from source run:

```shell
git clone https://github.com/goreapers/gograveyard
cd gograveyard
make
```

The default make command will run the required `go build` command and produce a
`gograveyard` binary in the root directory.

