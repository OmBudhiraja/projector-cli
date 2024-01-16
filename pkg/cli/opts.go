package cli

import (
	"flag"
	"fmt"
)

type Opts struct {
	Config string
	Pwd    string
	Args   []string
}

func GetOpts() *Opts {
	setupUsage()

	config := flag.String("c", "", "path to config file")
	pwd := flag.String("p", "", "path to project directory")
	flag.Parse()

	return &Opts{
		Config: *config,
		Pwd:    *pwd,
		Args:   flag.Args(),
	}
}

func setupUsage() {
	const usage = `usage: projector [options] [arguments]

Positional arguments:
  print
  add [key] [value]
  remove [key]

Options:
`

	flag.Usage = func() {
		fmt.Print(usage)
		flag.PrintDefaults()
		fmt.Println()
	}
}
