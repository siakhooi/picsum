/*
Package arguments to process arguments
*/
package arguments

import (
	"fmt"

	"github.com/siakhooi/picsum/internal/download"
	"github.com/siakhooi/picsum/internal/output"
	"github.com/siakhooi/picsum/internal/urlbuilder"
)

// Options holds all command-line flag values
type Options struct {
	ImageID    string
	Seed       string
	Grayscale  bool
	Blur       bool
	BlurLevel  int
	Quiet      bool
	OutputPath string
	Force      bool
}

// ValidateArguments validates the number of command-line arguments
func ValidateArguments(args []string) error {
	if len(args) == 0 || len(args) > 2 {
		return fmt.Errorf("invalid arguments")
	}
	return nil
}

// ValidateOptions validates flag values and applies business rules
func ValidateOptions(opts *Options) error {
	// Validate blur level range
	if opts.BlurLevel != 0 && (opts.BlurLevel < 1 || opts.BlurLevel > 10) {
		return fmt.Errorf("blur level must be between 1 and 10, got %d", opts.BlurLevel)
	}

	// If blurlevel is specified, it supersedes blur
	if opts.BlurLevel > 0 {
		opts.Blur = false
	}

	// Check mutual exclusivity
	if opts.ImageID != "" && opts.Seed != "" {
		return fmt.Errorf("options --id and --seed are mutually exclusive")
	}
	return nil
}

// ProcessImage handles the complete image processing workflow
func ProcessImage(args []string, opts *Options) error {
	// Build URL and filename based on arguments
	url, filename, err := urlbuilder.BuildURL(args, opts.ImageID, opts.Seed, opts.Grayscale, opts.Blur, opts.BlurLevel)
	if err != nil {
		return err
	}

	// Use custom output path if specified
	if opts.OutputPath != "" {
		filename = opts.OutputPath
	}

	// Download the image
	resp, err := download.Image(url, opts.Quiet)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	// Save the image to file
	return output.SaveImage(resp, filename, opts.Quiet, opts.Force)
}
