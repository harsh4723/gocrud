package main

import (
	"accountservice/cmd/flags"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	err := (&cli.App{
		Name:     "gocrud",
		Usage:    "Simple crud apis",
		Version:  "1.0",
		Before:   badge,
		Flags:    flags.Flags(),
		Commands: commands,
	}).Run(os.Args)
	if err != nil {
		fmt.Println("Something Went Wrong. Failed to start Gocrud.: " + err.Error())
		log.Fatal(
			fmt.Sprintf(
				"-- \nfailed to start Gocrud. \n--\n Caused By:\n%s\n--",
				errorstack(err.Error()),
			),
		)
	}
}
