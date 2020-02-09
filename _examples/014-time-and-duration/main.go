package main

import (
	"os"

	"github.com/franchb/cli"
	clix "github.com/franchb/cli/ext"
)

type argT struct {
	Time     clix.Time     `cli:"t" usage:"time"`
	Duration clix.Duration `cli:"d" usage:"duration"`
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("time=%v, duration=%v\n", argv.Time, argv.Duration)
		return nil
	}))
}
