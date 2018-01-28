package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:        "calc",
			Usage:       "nisekoi calc [<github-org> | <github-org/repo>]",
			Description: "Calculate average landing PR times",
			Action: func(c *cli.Context) error {
				username := c.String("username")
				fmt.Println(fmt.Sprintf("%s", username))
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "username, u",
					Usage:  "If set, average times for `USERNAME` will be displayed",
					EnvVar: "NISEKOI_USERNAME",
				},
			},
		},
	}
	app.Run(os.Args)
}
