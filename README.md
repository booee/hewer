# hewer
v0.2.0

### What is hewer?

hewer is a log file miner, a CLI app that can quickly parse and summarize large amounts of data from one or more log files.

The current implementation assumes a log file consists only of JSON content (one JSON blob per line).
From this, hewer will auto-discover JSON property keys that can be targeted for statistical analysis.
For details on the CLI, see [Usage](#usage) below.

The name "hewer" comes from [a term for a skilled miner](http://en.wikipedia.org/wiki/Hewer).

And, no, this has nothing to do with mining Bitcoin. Sorry.

### Installation Prerequisites

Eventually, hewer will be distributed as an executable binary that will have no prerequisites. Until then, you can install it via [Go](https://golang.org/).

1. Have Go installed ([let me google that for you](http://lmgtfy.com/?q=golang+install+download))
1. Environment setup according to specs on [golang.org](https://golang.org/doc)
  * [Workspace](https://golang.org/doc/code.html#Workspaces)
  * [Environment variable(s)](https://golang.org/doc/code.html#GOPATH) - and don't forget to add `$GOPATH/bin` to your `PATH`

### Installation

Install package via go tool

```
$ go get github.com/NexTraq/hewer
```

Verify Installation

```
$ hewer -v
```

### Usage

```
$ hewer [-k <keyName>] <fileName> [<anotherFileName>...]
```

Run `hewer --help` for more information
