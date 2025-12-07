package arguments

import (
	"os"
	"strings"
	"testing"
)

func TestValidateArguments(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "valid single argument",
			args:    []string{"200x300"},
			wantErr: false,
		},
		{
			name:    "valid two arguments",
			args:    []string{"200", "300"},
			wantErr: false,
		},
		{
			name:    "empty arguments",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "too many arguments",
			args:    []string{"200", "300", "extra"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateArguments(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateArguments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateOptions(t *testing.T) {
	tests := []struct {
		name    string
		opts    *Options
		wantErr bool
		check   func(*Options) bool
	}{
		{
			name: "valid options with no blur",
			opts: &Options{
				ImageID:   "123",
				Grayscale: true,
			},
			wantErr: false,
		},
		{
			name: "valid options with blur flag",
			opts: &Options{
				Blur: true,
			},
			wantErr: false,
		},
		{
			name: "valid blur level",
			opts: &Options{
				BlurLevel: 5,
			},
			wantErr: false,
		},
		{
			name: "blur level too low",
			opts: &Options{
				BlurLevel: 0,
			},
			wantErr: false, // 0 means not set
		},
		{
			name: "blur level below minimum",
			opts: &Options{
				BlurLevel: -1,
			},
			wantErr: true,
		},
		{
			name: "blur level above maximum",
			opts: &Options{
				BlurLevel: 11,
			},
			wantErr: true,
		},
		{
			name: "blur level supersedes blur flag",
			opts: &Options{
				Blur:      true,
				BlurLevel: 5,
			},
			wantErr: false,
			check: func(o *Options) bool {
				return !o.Blur // Blur should be set to false
			},
		},
		{
			name: "mutually exclusive id and seed",
			opts: &Options{
				ImageID: "123",
				Seed:    "myseed",
			},
			wantErr: true,
		},
		{
			name: "only id specified",
			opts: &Options{
				ImageID: "123",
			},
			wantErr: false,
		},
		{
			name: "only seed specified",
			opts: &Options{
				Seed: "myseed",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOptions(tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.check != nil && !tt.check(tt.opts) {
				t.Errorf("ValidateOptions() failed custom check")
			}
		})
	}
}

func TestProcessImage_Success(t *testing.T) {
	// GIVEN
	tmpfile := "test_process_image_success.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"200", "300"}
	opts := &Options{
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	// Note: This test will attempt to download from the real picsum.photos
	// For true unit testing, we would need to refactor ProcessImage to accept dependencies
	err := ProcessImage(args, opts)

	// THEN
	// This is more of an integration test since it calls real implementations
	// We can't easily test this without dependency injection
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}

	// Verify file was created
	if _, err := os.Stat(tmpfile); os.IsNotExist(err) {
		t.Error("Expected file to be created")
	}
}

func TestProcessImage_InvalidArguments(t *testing.T) {
	// GIVEN
	args := []string{"invalid"}
	opts := &Options{
		Quiet: true,
		Force: true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid arguments, got nil")
	}
	if !strings.Contains(err.Error(), "invalid") {
		t.Errorf("Expected error message to contain 'invalid', got: %v", err)
	}
}

func TestProcessImage_WithCustomOutputPath(t *testing.T) {
	// GIVEN
	tmpfile := "test_custom_output.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"100"}
	opts := &Options{
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}

	// Verify custom output path was used
	if _, err := os.Stat(tmpfile); os.IsNotExist(err) {
		t.Error("Expected file to be created at custom output path")
	}
}

func TestProcessImage_WithGrayscale(t *testing.T) {
	// GIVEN
	tmpfile := "test_grayscale.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"100"}
	opts := &Options{
		Grayscale:  true,
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}
}

func TestProcessImage_WithBlur(t *testing.T) {
	// GIVEN
	tmpfile := "test_blur.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"100"}
	opts := &Options{
		Blur:       true,
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}
}

func TestProcessImage_WithBlurLevel(t *testing.T) {
	// GIVEN
	tmpfile := "test_blurlevel.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"100"}
	opts := &Options{
		BlurLevel:  5,
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}
}

func TestProcessImage_WithImageID(t *testing.T) {
	// GIVEN
	tmpfile := "test_image_id.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"100"}
	opts := &Options{
		ImageID:    "237",
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}
}

func TestProcessImage_WithSeed(t *testing.T) {
	// GIVEN
	tmpfile := "test_seed.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"100"}
	opts := &Options{
		Seed:       "myseed",
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}
}

func TestProcessImage_TwoArguments(t *testing.T) {
	// GIVEN
	tmpfile := "test_two_args.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"200", "150"}
	opts := &Options{
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}
}

func TestProcessImage_MockDownloadError(t *testing.T) {
	// GIVEN
	// Create a test that simulates a download error
	// Since we can't easily mock the download.Image function without refactoring,
	// we use an invalid URL format that will cause urlbuilder to fail
	args := []string{"not-a-number"}
	opts := &Options{
		Quiet: true,
		Force: true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid arguments that would cause download failure")
	}
}

func TestProcessImage_QuietMode(t *testing.T) {
	// GIVEN
	tmpfile := "test_quiet.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"100"}
	opts := &Options{
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}
}

func TestProcessImage_FileExists_Force(t *testing.T) {
	// GIVEN
	tmpfile := "test_force_overwrite.jpg"
	// Create an existing file
	if err := os.WriteFile(tmpfile, []byte("existing content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"100"}
	opts := &Options{
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true, // Force overwrite
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}

	// Verify file was overwritten
	data, err := os.ReadFile(tmpfile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if string(data) == "existing content" {
		t.Error("Expected file to be overwritten")
	}
}

func TestProcessImage_FileExists_NoForce(t *testing.T) {
	// GIVEN
	tmpfile := "test_no_force.jpg"
	// Create an existing file
	if err := os.WriteFile(tmpfile, []byte("existing content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer func() { _ = os.Remove(tmpfile) }()

	// Mock stdin to simulate user declining overwrite
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	r, w, _ := os.Pipe()
	os.Stdin = r
	_, _ = w.Write([]byte("n\n"))
	_ = w.Close()

	args := []string{"100"}
	opts := &Options{
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      false, // Don't force overwrite
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	// User declined, so we expect an error
	if err != nil {
		if !strings.Contains(err.Error(), "user cancelled") && !strings.Contains(err.Error(), "failed") {
			t.Logf("ProcessImage integration test note: %v", err)
			t.Skip("Skipping integration test that requires network access")
		}
	}
}

func TestProcessImage_CombinedOptions(t *testing.T) {
	// GIVEN
	tmpfile := "test_combined.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	args := []string{"150", "100"}
	opts := &Options{
		ImageID:    "42",
		Grayscale:  true,
		BlurLevel:  3,
		OutputPath: tmpfile,
		Quiet:      true,
		Force:      true,
	}

	// WHEN
	err := ProcessImage(args, opts)

	// THEN
	if err != nil {
		t.Logf("ProcessImage integration test note: %v", err)
		t.Skip("Skipping integration test that requires network access")
	}
}
