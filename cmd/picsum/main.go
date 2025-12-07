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
	if err := cli.BuildCommand().Run(context.Background(), os.Args); err != nil {
		console.Stderrln("Error: %v", err)
		os.Exit(1)
	}
}
