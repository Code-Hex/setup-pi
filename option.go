package setup

import (
	"bytes"
	"fmt"
	"os"

	"reflect"

	"github.com/Code-Hex/exit"
	flags "github.com/jessevdk/go-flags"
)

const indent = "        "

// Options struct for parse command line arguments
type Options struct {
	Help    bool `short:"h" long:"help" description:"show this message"`
	Version bool `short:"v" long:"version" description:"print the version"`

	Filename   string `short:"f" long:"file" required:"true" description:"specify the yaml file"`
	StackTrace bool   `long:"trace" description:"display detail error messages"`
}

func (opts *Options) parse(argv []string) ([]string, error) {
	p := flags.NewParser(opts, flags.None)
	args, err := p.ParseArgs(argv)
	if err != nil {
		os.Stderr.Write(opts.usage())
		return nil, exit.MakeDataErr(err)
	}
	return args, nil
}

func (opts Options) usage() []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, `%s
Usage: %s [options]
Options:
`, msg, name)

	t := reflect.TypeOf(opts)
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag
		desc := tag.Get("description")
		var o string
		if s := tag.Get("short"); s != "" {
			o = fmt.Sprintf("-%s, --%s", tag.Get("short"), tag.Get("long"))
		} else {
			o = fmt.Sprintf("--%s", tag.Get("long"))
		}
		fmt.Fprintf(&buf, "  %-21s %s\n", o, desc)

		if deflt := tag.Get("default"); deflt != "" {
			fmt.Fprintf(&buf, "  %-21s default: --%s='%s'\n", indent, tag.Get("long"), deflt)
		}
	}

	return buf.Bytes()
}
