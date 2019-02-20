package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	flagHome = "home"
)

var (
	varHome *string

	// Version should be set by build flags: `git describe --tags`
	Version = "please set in makefile"
)

func init() {
	defaultHome := filepath.Join(os.ExpandEnv("$HOME"), ".tallybox")
	varHome = flag.String(flagHome, defaultHome, "directory to store files under")

	flag.CommandLine.Usage = helpMessage
}

func helpMessage() {
	fmt.Println("tallybox")
	fmt.Println("          Api server to accept votes and write to the blockchain")
	fmt.Println("")
	fmt.Println("help      Print this message")
	fmt.Println("keys      Creates keys if not existing, and lists addresses")
	fmt.Println("start     Run the api server")
	fmt.Println("version   Print software version")
	fmt.Println(`
  -home string
        directory to store files under (default "$HOME/.tallybox")`)
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Missing command:")
		helpMessage()
		os.Exit(1)
	}

	cmd := flag.Arg(0)
	rest := flag.Args()[1:]

	var err error
	switch cmd {
	case "help":
		helpMessage()
	case "keys":
		err = KeysCmd(*varHome)
	case "start":
		err = StartCmd(*varHome, rest)
	case "version":
		fmt.Println(Version)
	default:
		err = fmt.Errorf("unknown command: %s", cmd)
	}

	if err != nil {
		// fmt.Printf("Error: %+v\n\n", err)
		fmt.Printf("Error: %v\n\n", err)
		helpMessage()
		os.Exit(1)
	}
}
