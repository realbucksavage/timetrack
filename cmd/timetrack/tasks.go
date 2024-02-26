package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/realbucksavage/timetrack"
	"github.com/urfave/cli/v2"
)

var tasksCommand = &cli.Command{
	Name:    "task",
	Aliases: []string{"tasks", "t"},
	Subcommands: []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Action: func(cCtx *cli.Context) error {
				tasks, err := tracker.ListTasks()
				if err != nil {
					return err
				}

				slices.SortFunc(tasks, func(a, b *timetrack.Task) int {
					return strings.Compare(a.Bucket, b.Bucket)
				})

				for _, task := range tasks {
					fmt.Printf("%v\n", task)
				}

				return nil
			},
		},
		{
			Name:      "add",
			Aliases:   []string{"a"},
			Args:      true,
			ArgsUsage: "[bucket] [task] [duration]",
			Action: func(cCtx *cli.Context) error {
				args := cCtx.Args()
				if args.Len() != 3 {
					cli.ShowSubcommandHelpAndExit(cCtx, 1)
				}

				bucket := args.Get(0)
				task := args.Get(1)
				duration, err := time.ParseDuration(args.Get(2))
				if err != nil {
					return err
				}

				t, err := tracker.Track(bucket, task, duration)
				if err != nil {
					return err
				}

				fmt.Println(t)
				return nil
			},
		},
	},
}
