package setup

import (
	"fmt"
	"os"
)

// Setup context
type Setup struct{}

// New is construct
func New() *Setup {
	return &Setup{}
}

// Run command line
func (s *Setup) Run() int {
	if e := s.run(); e != nil {
		exitCode, err := UnwrapErrors(e)
		fmt.Fprintf(os.Stderr, "Error:\n  %v\n", err)
		return exitCode
	}
	return 0
}

func (s *Setup) run() error {
	return nil
}
