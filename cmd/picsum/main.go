/*
main cli entry
*/
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/siakhooi/picsum/internal/version"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "picsum",
		Usage:   "picsum application",
		Version: version.GetVersion(),
		Action: func(context.Context, *cli.Command) error {
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
