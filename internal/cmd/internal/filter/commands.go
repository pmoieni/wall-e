package filter

import (
	"github.com/disintegration/gift"
	"github.com/urfave/cli/v2"
)

var CtxName = "filter"

type FilterCtx struct {
	List map[string]gift.Filter
}

var Command = &cli.Command{
	Name: "filter",
	Subcommands: []*cli.Command{
		{
			Name:   "grayscale",
			Action: grayscale(),
		},
	},
}
