package main

import (
	"log"
	"os"
	"time"

	"github.com/pmoieni/wall-e/internal/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "wall-e",
		Usage:                "wallpaper engine for all kinds of linux desktops",
		EnableBashCompletion: true,
		Version:              cmd.Version,
		Compiled:             time.Now().UTC(),
		Action:               cmd.GetVersion,
		Commands:             cmd.GetCommands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
