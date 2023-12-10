package sources

import "github.com/urfave/cli/v2"

var Command = &cli.Command{
	Name:    "sources",
	Aliases: []string{"src"},
	Subcommands: []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "name",
					Aliases: []string{"n"},
					Usage:   "source name",
				},
				&cli.StringFlag{
					Name:    "type",
					Aliases: []string{"t"},
					Usage:   "source type",
				},
				&cli.StringFlag{
					Name:    "link",
					Aliases: []string{"l"},
					Usage:   "source link",
				},
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
		},
	},
}
