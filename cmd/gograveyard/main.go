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

func printHelp() {
	fmt.Println("The Go project undertaker: check go.mod dependency's health")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  gograveyard [flags] [command]")
	fmt.Println("")
	fmt.Println("Available Commands:")
	fmt.Println("  help     Print this help")
	fmt.Println("  parse    Parse a provided go.mod")
	fmt.Println("  version  Print current version")
	fmt.Println("")
	fmt.Println("Flags:")
	fmt.Println("  --help, -h   help for gograveyard")
	fmt.Println("  --json, -j   output final report in JSON")
}

func printVersion() {
	fmt.Printf("gograveyard (%s)\n", version)
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

	//nolint:godox
	// TODO: replace me with parsing the bytes
	fmt.Println(string(goModBytes))

	return nil
}

func main() {
	flag.Usage = printHelp

	var json bool
	flag.BoolVar(&json, "json", false, "output final report in JSON")
	flag.BoolVar(&json, "j", false, "output final report in JSON")

	// Parse the flags
	flag.Parse()

	// Parse the arguments, everything after the flags
	if len(os.Args) <= flag.NFlag()+1 {
		fmt.Printf("no subcommand specified")
		printHelp()
		os.Exit(1)
	}

	// First value is the subcommand, everything else are arguments passed to
	// the subcommand.
	subcommand := os.Args[flag.NFlag()+1]
	args := os.Args[flag.NFlag()+2:]

	switch subcommand {
	case "help":
		printHelp()
	case "parse":
		err := parse(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "version":
		printVersion()
	default:
		fmt.Printf("unknown subcommand")
		printHelp()
		os.Exit(1)
	}
}
