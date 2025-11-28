/*
main cli entry
*/
package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/siakhooi/picsum/internal/download"
	"github.com/siakhooi/picsum/internal/output"
	"github.com/siakhooi/picsum/internal/version"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "picsum",
		Usage:   "picsum <number> or picsum <width> <height>",
		Version: version.GetVersion(),
		Action: func(_ context.Context, c *cli.Command) error {
			args := c.Args().Slice()

			if len(args) == 0 || len(args) > 2 {
				return fmt.Errorf("invalid arguments\nUsage: picsum <number> or picsum <width> <height>")
			}

			var url, filename string

			if len(args) == 1 {
				// Parse single number
				num1, err := strconv.Atoi(args[0])
				if err != nil {
					return fmt.Errorf("invalid number: %s", args[0])
				}
				url = fmt.Sprintf("https://picsum.photos/%d", num1)
				filename = fmt.Sprintf("%d.jpg", num1)
			} else {
				// Parse two numbers
				num1, err := strconv.Atoi(args[0])
				if err != nil {
					return fmt.Errorf("invalid first number: %s", args[0])
				}
				num2, err := strconv.Atoi(args[1])
				if err != nil {
					return fmt.Errorf("invalid second number: %s", args[1])
				}
				url = fmt.Sprintf("https://picsum.photos/%d/%d", num1, num2)
				filename = fmt.Sprintf("%dx%d.jpg", num1, num2)
			}

			// Download the image
			resp, err := download.Image(url)
			if err != nil {
				return err
			}
			defer func() { _ = resp.Body.Close() }()

			// Save the image to file
			return output.SaveImage(resp, filename)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
