/*
main cli entry
*/
package main

import (
	"context"
	"os"

	"github.com/siakhooi/picsum/internal/cli"
	"github.com/siakhooi/picsum/internal/console"
)

func main() {
	if err := run(os.Args); err != nil {
		console.Stderrln("Error: %v", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	return cli.BuildCommand().Run(context.Background(), args)
}
