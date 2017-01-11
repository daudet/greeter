package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greeter"
	app.Usage = "Give a greeting!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello, there user!")
		return nil
	}

	app.Run(os.Args)
}
