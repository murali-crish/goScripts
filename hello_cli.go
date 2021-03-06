package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print hello world"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name, n",
			Usage: "Who to say hello to.",
			Value: "World",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s!\n", name)
		return nil
	}

	_ = app.Run(os.Args)
}
