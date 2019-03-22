package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type Options struct {
	showVersion bool
	array       bool
	pretty      bool
}

type Cli struct {
	inStream  io.Reader
	outStream io.Writer
	errStream io.Writer
	opts      *Options
}

const name = "gjo"
const version = "0.1.0"

var usage = fmt.Sprintf(`
%s - a copy of itchyny/gojo which is a go implementation of jo

version: %s

synopsis:
	$ gjo key1=value1 key1=value2 ...

options:
`, name, version)

const (
	exitCodeOK = iota
	exitCodeCliInitErr
)

func Run() int {
	c, err := newCli(os.Stdin, os.Stdout, os.Stderr, os.Args[1:])
	if err != nil {
		return exitCodeCliInitErr
	}

	return c.run()
}

func newCli(inStream io.Reader, outStream, errStream io.Writer, args []string) (*Cli, error) {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(errStream)
	fs.Usage = func() {
		_, _ = fmt.Fprintf(errStream, usage)
		fs.PrintDefaults()
	}
	var (
		showVersion bool
		array       bool
		pretty      bool
	)
	fs.BoolVar(&showVersion, "v", false, "show version")
	fs.BoolVar(&array, "a", false, "make value an array")
	fs.BoolVar(&pretty, "p", false, "pretty print")
	if err := fs.Parse(args); err != nil {
		return nil, err
	}
	options := &Options{showVersion, array, pretty}

	return &Cli{inStream, outStream, errStream, options}, nil
}

func (c *Cli) run() int {
	if c.opts.showVersion {
		fmt.Printf("%s - %s\n", name, version)
		return exitCodeOK
	}
	return exitCodeOK
}
