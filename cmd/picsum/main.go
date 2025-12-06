/*
main cli entry
*/
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/siakhooi/picsum/internal/cli"
)

func main() {
	if err := cli.BuildCommand().Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		os.Exit(1)
	}
}
