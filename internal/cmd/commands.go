package cmd

import (
	"github.com/pmoieni/wall-e/internal/cmd/internal/filter"
	"github.com/pmoieni/wall-e/internal/cmd/internal/scheduler"
	"github.com/pmoieni/wall-e/internal/cmd/internal/sources"
	"github.com/pmoieni/wall-e/internal/cmd/internal/store"
	"github.com/pmoieni/wall-e/internal/cmd/internal/wallpaper"
	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	cmds := []*cli.Command{}
	cmds = append(cmds,
		&cli.Command{
			Name:    "init",
			Aliases: []string{"i"},
		},
		filter.Command,
		scheduler.Command,
		sources.Command,
		store.Command,
		wallpaper.Command,
	)

	return cmds
}
