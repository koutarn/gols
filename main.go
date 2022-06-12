package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

const (
	appName    = "gols"
	appVersion = "0.01"
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErrArgs
)

type options struct {
	Version bool `short:"v" long:"version" description:"Show version"`
}

func main() {
	code, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err)
	}

	os.Exit(int(code))
}

func run(cliantArgs []string) (exitCode, error) {
	var opts options
	parser := flags.NewParser(&opts, flags.Default)

	args, err := parser.ParseArgs(cliantArgs)

	if err != nil {
		if flags.WroteHelp(err) {
			return exitCodeOK, nil
		}
		return exitCodeErrArgs, fmt.Errorf("parse error:%w", err)
	}

	if len(args) == 0 {
		return exitCodeErrArgs, errors.New("must requires an arguments")
	}

	if opts.Version {
		fmt.Printf("%s: v%s\n", appName, appVersion)
		return exitCodeOK, nil
	}

	return exitCodeOK, nil
}
