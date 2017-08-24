package setup

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/Code-Hex/exit"
	"github.com/k0kubun/pp"
)

//go:generate go-bindata -pkg setup -o bindata.go ./files/...

// Setup context
type Setup struct {
	Files  []string
	Recipe *Recipe
}

// New is construct
func New() *Setup {
	return &Setup{
		Files:  AssetNames(),
		Recipe: new(Recipe),
	}
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
	pp.Println(s.Files)
	f, err := os.Open("main.yml")
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

	pp.Println(s.Recipe)

	return nil
}
