package setup

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"

	"gopkg.in/yaml.v2"

	"github.com/Code-Hex/exit"
	"github.com/pkg/errors"
)

//go:generate go-bindata -pkg setup -o bindata.go ./files/...
const (
	version = "0.0.1"
	msg     = name + " v" + version + ", Setup tool for raspberry pi"
	name    = "setup"
)

// Setup context
type Setup struct {
	Options
	Recipe *Recipe
	Files  []string
	Env    []string
	Blue   *color.Color
	Yellow *color.Color
}

// New is construct
func New() *Setup {
	return &Setup{
		Recipe: new(Recipe),
		Files:  AssetNames(),
		Env:    os.Environ(),
		Blue:   color.New(color.FgBlue),
		Yellow: color.New(color.FgYellow),
	}
}

// Run command line
func (s *Setup) Run() int {
	if e := s.run(); e != nil {
		exitCode, err := UnwrapErrors(e)
		if s.StackTrace {
			fmt.Fprintf(os.Stderr, "Error:\n  %+v\n", e)
		} else {
			fmt.Fprintf(os.Stderr, "Error:\n  %v\n", err)
		}
		return exitCode
	}
	return 0
}

func (s *Setup) run() error {
	if os.Geteuid() > 0 {
		return errors.New("You must run this program as a superuser")
	}
	_, err := parseOptions(&s.Options, os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "Failed to parse command line args")
	}
	if err := s.parseRecipe(); err != nil {
		return errors.Wrap(err, "Failed to parse recipe")
	}
	if err := s.execute(); err != nil {
		return errors.Wrap(err, "Failed to execute item")
	}
	return nil
}

func parseOptions(opts *Options, argv []string) ([]string, error) {
	o, err := opts.parse(argv)
	if err != nil {
		return nil, exit.MakeDataErr(err)
	}
	if opts.Help {
		return nil, exit.MakeUsage(errors.New(string(opts.usage())))
	}
	if opts.Version {
		return nil, exit.MakeUsage(errors.New(msg))
	}
	return o, nil
}

func (s *Setup) parseRecipe() error {
	f, err := os.Open(s.Filename)
	if err != nil {
		return exit.MakeOSFile(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return exit.MakeSoftWare(err)
	}

	if err := yaml.Unmarshal(b, s.Recipe); err != nil {
		return exit.MakeDataErr(err)
	}
	return nil
}
