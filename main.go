package main

import (
	"encoding/json"
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

	projector := cli.NewProjector(config)

	if config.Operation == cli.Print {
		if len(config.Args) == 0 {
			data := projector.GetValueAll()
			jsonString, err := json.Marshal(data)

			if err != nil {
				log.Fatalf("this line should never be reached %v", err)
			}

			fmt.Println(string(jsonString))

		} else {
			value, ok := projector.GetValue(config.Args[0])
			if ok {
				fmt.Println(value)
			}
		}
	} else if config.Operation == cli.Add {
		projector.AddValue(config.Args[0], config.Args[1])
		projector.Save()
	} else {
		projector.RemoveValue(config.Args[0])
		projector.Save()
	}
}
