package main

import (
	"os"

	"github.com/franchb/cli"
	clix "github.com/franchb/cli/ext"
)

type argT struct {
	cli.Helper
	PidFile clix.PidFile `cli:"pid" usage:"pid file" dft:"013-pidfile.pid"`
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)

		if err := argv.PidFile.New(); err != nil {
			return err
		}
		defer argv.PidFile.Remove()

		return nil
	}))
}
