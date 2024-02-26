package main

import (
	"cmp"
	"fmt"
	"slices"
	"time"

	"github.com/realbucksavage/timetrack"
	"github.com/urfave/cli/v2"
)

var tasksCommand = &cli.Command{
	Name:    "task",
	Aliases: []string{"tasks", "t"},
	Subcommands: []*cli.Command{
		menuCommand,
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Action: func(cCtx *cli.Context) error {
				tasks, err := tracker.ListTasks()
				if err != nil {
					return err
				}

				for bucket, bucketTasks := range tasks {
					slices.SortFunc(bucketTasks, func(a, b *timetrack.Task) int {
						return cmp.Compare(a.Spent, b.Spent)
					})

					var (
						sum     time.Duration
						taskStr string
					)

					for _, t := range bucketTasks {
						sum += t.Spent
						taskStr = fmt.Sprintf("%s\t(%v) %s\n", taskStr, t.Spent, t.Task)
					}

					fmt.Printf("%s: %v\n%s", bucket, sum, taskStr)
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
