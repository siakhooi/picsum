/*
main cli entry
*/
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/siakhooi/picsum/internal/download"
	"github.com/siakhooi/picsum/internal/output"
	"github.com/siakhooi/picsum/internal/urlbuilder"
	"github.com/siakhooi/picsum/internal/version"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "picsum",
		Usage:   "fetch photo from https://picsum.photos",
		Version: version.Version(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "id",
				Aliases: []string{"i"},
				Usage:   "specific image ID from picsum.photos",
			},
			&cli.StringFlag{
				Name:    "seed",
				Aliases: []string{"s"},
				Usage:   "seed for random image generation from picsum.photos",
			},
			&cli.BoolFlag{
				Name:    "gray",
				Aliases: []string{"g"},
				Usage:   "convert image to grayscale",
			},
		},
		Action: func(_ context.Context, c *cli.Command) error {
			args := c.Args().Slice()

			if len(args) == 0 || len(args) > 2 {
				return fmt.Errorf("invalid arguments\nUsage: picsum [-h] [-v] [-i imageId | -s seed] (size|width height)")
			}

			imageID := c.String("id")
			seed := c.String("seed")
			grayscale := c.Bool("gray")

			// Check mutual exclusivity
			if imageID != "" && seed != "" {
				return fmt.Errorf("options --id and --seed are mutually exclusive")
			}
			// Build URL and filename based on arguments
			url, filename, err := urlbuilder.BuildURL(args, imageID, seed, grayscale)
			if err != nil {
				return err
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
