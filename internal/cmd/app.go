package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const Version = "v0.0.0-a.1"

func GetVersion(cCtx *cli.Context) error {
	_, err := fmt.Println("wall-e version: " + Version)
	return err
}
