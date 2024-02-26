package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var bucketsCommand = &cli.Command{
	Name:    "bucket",
	Aliases: []string{"buckets", "b"},
	Subcommands: []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Action: func(cCtx *cli.Context) error {
				buckets, err := tracker.ListBuckets()
				if err != nil {
					return err
				}

				for _, b := range buckets {
					fmt.Println(b)
				}

				return nil
			},
		},
	},
}
