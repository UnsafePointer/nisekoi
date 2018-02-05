package main

import (
	"os"

	"github.com/Ruenzuo/nisekoi/calc"
	"github.com/Ruenzuo/nisekoi/utils"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:        "calc",
			Usage:       "nisekoi calc [<owner> | <owner/repo>]",
			Description: "Calculate average landing PR times",
			Action: func(c *cli.Context) error {
				lookup := c.Args().First()
				owner, repo, err := utils.ValidateSearchTerm(lookup)
				if err != nil {
					return cli.NewExitError("The search term doesn't conform to [<owner> | <owner/repo>]", 1)
				}

				username := c.String("username")
				if len(username) > 0 {
					if utils.ValidateIdentifier(username) != nil {
						return cli.NewExitError("The username provided is invalid", 2)
					}
				}

				err = calc.Cmd{
					Owner:       owner,
					Repository:  repo,
					Username:    username,
					AccessToken: c.String("access-token"),
					Debug:       c.Bool("debug"),
				}.Run()
				if err != nil {
					return cli.NewExitError(err.Error(), 3)
				}
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "username, u",
					Usage:  "If set, average times for `USERNAME` will be displayed",
					EnvVar: "NISEKOI_USERNAME",
				},
				cli.StringFlag{
					Name:   "access-token, t",
					Usage:  "If set, Nisekoi will use the access token provided for authentication",
					EnvVar: "NISEKOI_ACCESS_TOKEN",
				},
				cli.BoolFlag{
					Name:   "debug",
					Usage:  "If set, debug output will be printed",
					EnvVar: "NISEKOI_DEBUG",
				},
			},
		},
	}
	app.Run(os.Args)
}
