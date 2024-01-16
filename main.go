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

	fmt.Println(config)

	/*	var file *os.File
		if _, err := os.Stat(config.Config); os.IsNotExist(err) {
			fmt.Println("creating config file")
			os.MkdirAll(filepath.Dir(config.Config), os.ModePerm)
			os.Create(config.Config)
		} else {
			file, _ = os.OpenFile(config.Config, os.O_RDWR, 0644)
		}

		data, err := os.ReadFile(config.Config)

		if err != nil {
			log.Fatalf("could not create config file %v", err)
		}

		_, err = os.Stat(config.Pwd)

		if err != nil {
			log.Fatal("could not find specified directory")
		}

		fmt.Println(string(data), data)

		if err != nil {
			log.Fatal("could not read the config file")
		}

		store := make(map[string]map[string]string)

		err = json.Unmarshal(data, &store)

		if err != nil {
			log.Fatalf("could not unmarshal the config file %v", err)
		}

		if config.Operation == cli.Add {
			// create a map of string and strings
			store[config.Pwd] = map[string]string{}
			store[config.Pwd][config.Args[0]] = config.Args[1]
		} else if config.Operation == cli.Remove {
			delete(store[config.Pwd], config.Args[0])
		} else {
			fmt.Println(store[config.Pwd])
			return
		}

		d, err := json.Marshal(store)

		if err != nil {
			log.Fatalf("error encoding json %v", err)
		}

		if err != nil {
			log.Fatalf("error opening file %v", err)
		}

		err = file.Truncate(0)
		_, err = file.Seek(0, 0)
		fmt.Println(store)
		_, err = file.Write(d)

		if err != nil {
			log.Fatalf("error writing to file %v", err)
		}
	*/
}
