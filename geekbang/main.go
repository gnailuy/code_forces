package main

import (
	"flag"
	"fmt"
	"os"
)

// var cmdLine = flag.NewFlagSet("question", flag.PanicOnError)
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)
var name string

func init() {
	cmdLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of question:\n")
		cmdLine.PrintDefaults()
	}
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
