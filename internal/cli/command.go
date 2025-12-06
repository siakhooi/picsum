/*
Package cli to build and run cli command
*/
package cli

import (
	"context"

	"github.com/siakhooi/picsum/internal/arguments"
	"github.com/siakhooi/picsum/internal/version"
	"github.com/urfave/cli/v3"
)

// BuildCommand creates and configures the CLI command
func BuildCommand() *cli.Command {
	return &cli.Command{
		Name:    "picsum",
		Usage:   "fetch photo from https://picsum.photos",
		Version: version.Version(),
		Flags:   buildFlags(),
		Action:  runAction,
	}
}

func buildFlags() []cli.Flag {
	return []cli.Flag{
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
		&cli.BoolFlag{
			Name:    "blur",
			Aliases: []string{"b"},
			Usage:   "apply blur effect to image",
		},
		&cli.IntFlag{
			Name:    "blurlevel",
			Aliases: []string{"B"},
			Usage:   "apply blur effect with specific level 1-10 (supersedes -b)",
		},
		&cli.BoolFlag{
			Name:    "quiet",
			Aliases: []string{"q"},
			Usage:   "suppress output messages",
		},
	}
}

func runAction(_ context.Context, c *cli.Command) error {
	args := c.Args().Slice()

	if err := arguments.ValidateArguments(args); err != nil {
		return err
	}

	opts := &arguments.Options{
		ImageID:   c.String("id"),
		Seed:      c.String("seed"),
		Grayscale: c.Bool("gray"),
		Blur:      c.Bool("blur"),
		BlurLevel: c.Int("blurlevel"),
		Quiet:     c.Bool("quiet"),
	}

	if err := arguments.ValidateOptions(opts); err != nil {
		return err
	}

	return arguments.ProcessImage(args, opts)
}
