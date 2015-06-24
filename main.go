package main

import (
	"fmt"
	"github.com/mgutz/minimist"
	"os"
)

var name = "hewer"
var version = "0.2.0"

var usage = `
Mine the crap out of some JSON-based log files

Usage: hewer [options] file [files...]
Options:
   -k <key>, --key <key>        Operate on this JSON key, either collecting
                                stats or converting its values

   -c <type>, --convert <type>  Convert the value at <key> to this type
                                and print modified JSON to stdout. Valid
                                values are number, json, string

   -h, --help                   Print this content and exit

   -v, --version                Print version and exit
`

func main() {
	// parse command line args
	argm := minimist.Parse()

	// check for help flag
	if argm.AsBool("help", "h", "?") {
		helpAndExit()
	}

	// check for version flag
	if argm.AsBool("v", "version") {
		fmt.Println(version)
		os.Exit(0)
	}

	// check for file args
	args := argm["_"].([]string)
	if len(args) < 1 {
		helpAndExit()
	}

	//TODO verify existence of files

	// grab optional "key" flag and init analytics
	key := argm.MayString("", "key", "k")

	convert := argm.MayString("", "convert", "c")

	//TODO parse mulitple files concurrently
	if convert == "" {
		analytics := NewAnalytics(key)
		for _, value := range args {
			ParseFile(value, analytics)
		}
		// print all results
		analytics.Print()
	} else {
		converter := NewConverter(key, convert)
		for _, value := range args {
			ParseAndConvertFile(value, converter)
		}
	}
}

func helpAndExit() {
	fmt.Println(usage)
	os.Exit(0)
}
