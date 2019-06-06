// Copyright (c) 2019 Luiz Poleto
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lpoleto/serier/episodefile"
)

// VERSION holds the current version.
// The actual version will be provided during build using the flag: -ldflags "-X main.VERSION=X.Y.Z"
var VERSION = "UNKNOWN"

// Options represents the command line options provided when calling the program.
type Options struct {
	SeriesName string
	Dir        string
}

// Initialize the command line options and check if the required parameters were provided.
// Returns an Options struct with the parsed options.
//
func initArgs() *Options {
	options := &Options{}

	flag.StringVar(&options.SeriesName, "s", "", "Series name. It will be used as part of the filename and also for lookup on TV Database.")
	flag.StringVar(&options.Dir, "p", ".", "Path to scan for series episodes.")
	versionFlag := flag.Bool("version", false, "Print version")

	flag.Parse()

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}

	if options.SeriesName == "" {
		printVersion()
		fmt.Println("Series name must be provided. Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	return options
}

func printVersion() {
	fmt.Printf("Serier version %s\n", VERSION)
}

func main() {
	options := initArgs()

	if !episodefile.FileExists(options.Dir) {
		fmt.Printf("The path provided [%s] does not exists. Aborting.\n", options.Dir)
		os.Exit(-1)
	}

	err := ReadSeries(options.SeriesName, options.Dir)

	if err != nil {
		fmt.Printf("An error occurred: %v\n\n", err)
	}
}
