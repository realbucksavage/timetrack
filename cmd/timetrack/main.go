package main

import (
	"os"

	"github.com/realbucksavage/timetrack"
	"github.com/urfave/cli/v2"
)

func main() {

	var (
		baseDir string
		tracker *timetrack.Tracker
	)

	app := &cli.App{
		Name:        "timetrack",
		Description: "A simple utility to track time spent over various `buckets` of tasks",
		Version:     "0.1.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "base-dir",
				Aliases:     []string{"d"},
				Usage:       "The base directory where `timetrack` will store it's data.",
				Destination: &baseDir,
			},
		},
		Before: func(cCtx *cli.Context) error {
			var err error

			options := make([]timetrack.Option, 0)
			if baseDir != "" {
				options = append(options, timetrack.WithBaseDir(baseDir))
			}

			tracker, err = timetrack.NewTracker(options...)
			return err
		},
		Commands: []*cli.Command{
			{
				Name:  "status",
				Usage: "check for errors and such.",
				Action: func(cCtx *cli.Context) error {
					cli.ShowVersion(cCtx)
					return tracker.Status()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
