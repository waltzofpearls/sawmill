package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sawmill"
	app.Usage = "Look up possible malware infected URLs"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}
	app.Run(os.Args)
}
