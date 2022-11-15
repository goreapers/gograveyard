package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/goreapers/gograveyard"
)

const (
	version string = "v0.0.1"
)

var (
	errNoFileArg    = errors.New("a path to go.mod is required")
	errTooManyFiles = errors.New("too many arguments, only one go.mod allowed")
)

func helpString() string {
	return `The Go project undertaker: check go.mod dependency's health

Usage:
  gograveyard [flags] [command]

Available Commands:
  help     Print this help
  parse    Parse a provided go.mod
  version  Print current version

Flags:
  --help, -h   help for gograveyard
  --json, -j   output final report in JSON
`
}

func versionString() string {
	return fmt.Sprintf("gograveyard (%s)\n", version)
}

func parse(args []string) error {
	if len(args) == 0 {
		return errNoFileArg
	}
	if len(args) > 1 {
		return errTooManyFiles
	}

	goModBytes, err := gograveyard.GoModBytes(args[0])
	if err != nil {
		return fmt.Errorf("unable to read '%s': %w", args[0], err)
	}

	modFile, err := gograveyard.Parse(goModBytes)
	if err != nil {
		return fmt.Errorf("failed to parse: %w", err)
	}

	fmt.Printf("This modfile has %d direct dependencies and %d indirect dependencies \n",
		len(modFile.Direct), len(modFile.Indirect))

	return nil
}

func main() {
	flag.Usage = func() {
		fmt.Print(helpString())
	}

	var json bool
	flag.BoolVar(&json, "json", false, "output final report in JSON")
	flag.BoolVar(&json, "j", false, "output final report in JSON")

	// Parse the flags
	flag.Parse()

	// Parse the arguments, everything after the flags
	if len(os.Args) <= flag.NFlag()+1 {
		fmt.Print("no subcommand specified\n")
		fmt.Print(helpString())
		os.Exit(1)
	}

	// First value is the subcommand, everything else are arguments passed to
	// the subcommand.
	subcommand := os.Args[flag.NFlag()+1]
	args := os.Args[flag.NFlag()+2:]

	switch subcommand {
	case "help":
		fmt.Print(helpString())
	case "parse":
		err := parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "version":
		fmt.Print(versionString())
	default:
		fmt.Printf("unknown subcommand: '%s'\n", subcommand)
		fmt.Print(helpString())
		os.Exit(1)
	}
}
