# hewer
v0.2.0

### What is hewer?

hewer is a log file miner, a CLI app that can quickly parse and summarize large amounts of data from one or more log files.

The current implementation assumes a log file consists of JSON content (one JSON blob per line).
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

#### Summarize JSON data

```
$ hewer [-k <keyName>] <fileName> [<anotherFileName>...]
$ hewer <fileName> [<anotherFileName>...] [-k <keyName>]
```

Collect stats for a given JSON key (possibly nested) and print the collected info to stdout. If no key is given, hewer will default to the root object of each JSON blob. Can be used to "discover" the JSON structure(s) within the file(s).

#### Convert JSON values

```
$ hewer -k <keyName> -c <type> <fileName> [<anotherFileName>...]
$ hewer <fileName> [<anotherFileName>...] -k <keyName> -c <type>
```

Convert the values for a given JSON key to a more useful type and print the modified JSON to stdout, which you can redirect to write to another file. Valid values for `<type>` are:

- `number` (convert string to floating point or integer number)
- `json` (convert a JSON-encoded string to raw JSON)
- `string` (convert something to a string).

Lines from the input file(s) that are not JSON will be preserved as is. If no key is given, no conversion will be done and the original contents of the file(s) will be echoed.

Run `hewer --help` for more information
