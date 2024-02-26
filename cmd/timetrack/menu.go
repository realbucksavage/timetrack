package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// menuCommand serves as a potential starting place for integrating with systems like dmenu or rofi.
var menuCommand = &cli.Command{
	Name: "menu",
	Action: func(cCtx *cli.Context) error {
		lineage := cCtx.Lineage()
		if len(lineage) > 2 {
			cCtx = lineage[1]
		}

		commands := cCtx.Command.VisibleCommands()
		for _, cmd := range commands {
			if cmd.Name == "menu" || cmd.Name == "help" {
				continue
			}

			fmt.Println(cmd.FullName())
		}

		return nil
	},
}
