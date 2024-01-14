package main

import (
	"fmt"
	"log"

	"github.com/ombudhiraja/projector-cli/pkg/cli"
)

func main() {
	opts := cli.GetOpts()

	config, err := cli.NewConfig(opts)

	if err != nil {
		log.Fatalf("unable to get config: %v", err)
	}

	fmt.Printf("Opts: %+v\n", opts)
	fmt.Printf("Config: %+v\n", config)

}
