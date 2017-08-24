package main

import (
	"os"

	setup "github.com/Code-Hex/setup-pi"
)

func main() {
	os.Exit(setup.New().Run())
}
