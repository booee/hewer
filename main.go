package main

import (
	"fmt"
	"github.com/mgutz/minimist"
	"os"
)

var name = "hewer"
var version = "0.1.0"

var usage = `
hewer - Mine the crap out of some JSON-only log files

Usage: hewer [options] file [files...]
Options:
   -k <key>, --key <key>  Collect stats for this JSON key
   -h,       --help       Print this content and exit
   -v,       --version    Print version and exit
`

func main() {
	// ParseFile("/Users/bradbowie/Desktop/metrics-calamp-blue-20501.log");

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

	// fmt.Println("Mining files:", args)

	// grab optional "key" flag and init analytics
	key := argm.MayString("", "key", "k")
	// fmt.Println("Using key:", key)

	//analytics := NewAnalytics("system.activeThreads")
	analytics := NewAnalytics(key)

	//TODO parse mulitple files concurrently
	for _, value := range args {
		//fmt.Println(value)
		ParseFile(value, analytics)
	}

	// print all results
	analytics.Print()
}

func helpAndExit() {
	fmt.Println(usage)
	os.Exit(0)
}
