package main

import (
	"nisekoi/calc"
	"nisekoi/utils"
	"os"

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

				calc.Cmd{
					Owner:      owner,
					Repository: repo,
					Username:   username,
				}.Run()
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
