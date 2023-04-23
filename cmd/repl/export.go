package repl

import "github.com/urfave/cli/v2"

var ExportCmd = &cli.Command{
	Name:      "export",
	Usage:     "Check the status of a replication request",
	ArgsUsage: "REQUEST_ID",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "watch",
			Usage:   "Watch the progress of the replication",
			Aliases: []string{"w"},
			Value:   false,
		},
	},
	Action: func(c *cli.Context) error {
		return nil
	},
}
