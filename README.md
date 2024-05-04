# Codeforces Tool CP Editor edition


[![Github release](https://img.shields.io/github/release/aqua4/cf-tool-cpe.svg)](https://github.com/aqua4/cf-tool-cpe/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/aqua4/cf-tool-cpe)](https://goreportcard.com/report/github.com/aqua4/cf-tool-cpe)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-green.svg)](https://github.com/golang)
[![license](https://img.shields.io/badge/license-MIT-%23373737.svg)](https://raw.githubusercontent.com/aqua4/cf-tool-cpe/main/LICENSE)

Codeforces Tool CP Editor edition is a command-line interface tool for [Codeforces](https://codeforces.com) and [CP Editor](https://github.com/cpeditor/cpeditor/).

It's fast, small and cross-platform.

Simplified version of [Original CF tool](https://github.com/xalanq/cf-tool).

[Installation](#installation) | [Usage](#usage) | [FAQ](#faq)

## Features

* Supports Contests and Gym.
* Supports C++20 in Codeforces.
* Submit codes.
* Watch a submission's status dynamically.
* Setup a network proxy. Setup a mirror host.

Pull requests are always welcome.

## Installation

You can download the pre-compiled binary file in [here](https://github.com/aqua4/cf-tool-cpe/releases).

Or you can compile it from the source **(go >= 1.17)**:

```plain
$ go get github.com/aqua4/cf-tool-cpe
$ cd $GOPATH/.../cf-tool-cpe
$ go build -ldflags "-s -w" cf.go
```

If you don't know what's the `$GOPATH`, please see here <https://github.com/golang/go/wiki/GOPATH>.

## Usage
```plain
You should run "cf config" to configure your handle and password at first.
Usage:
  cf config
  cf submit -f <file> <url>
Options:
  -h --help     Show this screen.
  --version     Show version.
  -f <file>, --file <file>
                Path to file. E.g. "a.cpp", "./temp/a.cpp"
  <url>         Problem URL. E.g. "https://codeforces.com/contest/180/problem/A",
Examples:
  cf config            Configure the cf-tool.
  cf submit -f a.cpp https://codeforces.com/contest/100/A
File:
  cf will save some data in some files:
  "~/.cf/config"        Configuration file.
  "~/.cf/session"       Session file, including cookies, handle, password, etc.
  "~" is the home directory of current user in your system.
```

## Main changes

* Removed most of the features which are already supported by CP editor like templates, colors, etc.
* Removed most of the CLI commands which are not used by CP editor.
* Removed all programming languages support.
* Added C++20 support.

## FAQ

### I double click the program but it doesn't work

Codeforces Tool is a command-line tool. You should run it in terminal.

### I cannot use `cf` command

You should put the `cf` program to a path (e.g. `/usr/bin/` in Linux) which has been added to system environment variable PATH.

Or just google "how to add a path to system environment variable PATH".

### I want to change C++ compiler version or source file extension:

**Reminder:** C++20 is backwards compatible with previous versions, i.e. you can submit C++11/14/17 with C++20.

Change [supported languages](https://github.com/aqua4/cf-tool-cpe/blob/main/client/langs.go) and re-build the binary.
