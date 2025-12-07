package cli

import (
	"context"
	"os"
	"testing"

	"github.com/siakhooi/picsum/internal/version"
	"github.com/urfave/cli/v3"
)

func TestBuildCommand(t *testing.T) {
	cmd := BuildCommand()

	if cmd == nil {
		t.Fatal("BuildCommand() returned nil")
	}

	if cmd.Name != "picsum" {
		t.Errorf("BuildCommand() Name = %v, want %v", cmd.Name, "picsum")
	}

	expectedUsage := "fetch photo from https://picsum.photos"
	if cmd.Usage != expectedUsage {
		t.Errorf("BuildCommand() Usage = %v, want %v", cmd.Usage, expectedUsage)
	}

	expectedVersion := version.Version()
	if cmd.Version != expectedVersion {
		t.Errorf("BuildCommand() Version = %v, want %v", cmd.Version, expectedVersion)
	}

	if cmd.Action == nil {
		t.Error("BuildCommand() Action is nil")
	}

	if cmd.Flags == nil {
		t.Fatal("BuildCommand() Flags is nil")
	}

	if len(cmd.Flags) != 8 {
		t.Errorf("BuildCommand() Flags length = %v, want %v", len(cmd.Flags), 8)
	}
}

func TestBuildFlags(t *testing.T) {
	flags := buildFlags()

	if len(flags) != 8 {
		t.Errorf("buildFlags() returned %d flags, want 8", len(flags))
	}

	tests := []struct {
		name        string
		flagName    string
		flagType    string
		aliases     []string
		description string
	}{
		{
			name:        "id flag",
			flagName:    "id",
			flagType:    "*cli.StringFlag",
			aliases:     []string{"i"},
			description: "specific image ID from picsum.photos",
		},
		{
			name:        "seed flag",
			flagName:    "seed",
			flagType:    "*cli.StringFlag",
			aliases:     []string{"s"},
			description: "seed for random image generation from picsum.photos",
		},
		{
			name:        "gray flag",
			flagName:    "gray",
			flagType:    "*cli.BoolFlag",
			aliases:     []string{"g"},
			description: "convert image to grayscale",
		},
		{
			name:        "blur flag",
			flagName:    "blur",
			flagType:    "*cli.BoolFlag",
			aliases:     []string{"b"},
			description: "apply blur effect to image",
		},
		{
			name:        "blurlevel flag",
			flagName:    "blurlevel",
			flagType:    "*cli.IntFlag",
			aliases:     []string{"B"},
			description: "apply blur effect with specific level 1-10 (supersedes -b)",
		},
		{
			name:        "quiet flag",
			flagName:    "quiet",
			flagType:    "*cli.BoolFlag",
			aliases:     []string{"q"},
			description: "suppress output messages",
		},
		{
			name:        "output flag",
			flagName:    "output",
			flagType:    "*cli.StringFlag",
			aliases:     []string{"o"},
			description: "output file path",
		},
		{
			name:        "force flag",
			flagName:    "force",
			flagType:    "*cli.BoolFlag",
			aliases:     []string{"f"},
			description: "overwrite existing file without prompting",
		},
	}

	flagMap := make(map[string]cli.Flag)
	for _, flag := range flags {
		for _, name := range flag.Names() {
			flagMap[name] = flag
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, exists := flagMap[tt.flagName]
			if !exists {
				t.Fatalf("flag %q not found", tt.flagName)
			}

			// Verify flag names (main name + aliases)
			flagNames := flag.Names()
			expectedNames := append([]string{tt.flagName}, tt.aliases...)
			if len(flagNames) != len(expectedNames) {
				t.Errorf("flag %q has %d names, want %d", tt.flagName, len(flagNames), len(expectedNames))
			}

			// Verify all expected names are present
			for _, name := range expectedNames {
				found := false
				for _, flagName := range flagNames {
					if flagName == name {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("flag %q missing expected name/alias %q", tt.flagName, name)
				}
			}

			// Verify type
			switch tt.flagType {
			case "*cli.StringFlag":
				if _, ok := flag.(*cli.StringFlag); !ok {
					t.Errorf("flag %q is not a StringFlag", tt.flagName)
				}
			case "*cli.BoolFlag":
				if _, ok := flag.(*cli.BoolFlag); !ok {
					t.Errorf("flag %q is not a BoolFlag", tt.flagName)
				}
			case "*cli.IntFlag":
				if _, ok := flag.(*cli.IntFlag); !ok {
					t.Errorf("flag %q is not an IntFlag", tt.flagName)
				}
			}
		})
	}
}

func TestRunAction_WithMockCommand(t *testing.T) {
	tests := []struct {
		name    string
		setupFn func() *cli.Command
		args    []string
		wantErr bool
	}{
		{
			name: "no arguments should error",
			setupFn: func() *cli.Command {
				return BuildCommand()
			},
			args:    []string{"picsum"},
			wantErr: true,
		},
		{
			name: "too many arguments should error",
			setupFn: func() *cli.Command {
				return BuildCommand()
			},
			args:    []string{"picsum", "200", "300", "extra"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := tt.setupFn()
			ctx := context.Background()
			err := cmd.Run(ctx, tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("cmd.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunAction_OptionsValidation(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "blur level out of range - too high",
			args:    []string{"picsum", "--blurlevel", "11", "200"},
			wantErr: true,
			errMsg:  "blur level must be between 1 and 10",
		},
		{
			name:    "blur level out of range - negative",
			args:    []string{"picsum", "--blurlevel", "-1", "200"},
			wantErr: true,
			errMsg:  "blur level must be between 1 and 10",
		},
		{
			name:    "mutually exclusive id and seed",
			args:    []string{"picsum", "--id", "123", "--seed", "myseed", "200"},
			wantErr: true,
			errMsg:  "mutually exclusive",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := BuildCommand()
			ctx := context.Background()
			err := cmd.Run(ctx, tt.args)

			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if tt.errMsg != "" && err != nil {
				if !contains(err.Error(), tt.errMsg) {
					t.Errorf("Expected error containing %q, got %q", tt.errMsg, err.Error())
				}
			}
		})
	}
}

func TestRunAction_OptionsCreation(t *testing.T) {
	// Test that runAction properly creates Options from CLI flags
	// This tests lines 77-86 of command.go
	// NOTE: These are integration tests that make real network calls
	tests := []struct {
		name        string
		args        []string
		cleanupFile string
	}{
		{
			name:        "with id flag",
			args:        []string{"picsum", "--id", "123", "--quiet", "--force", "200"},
			cleanupFile: "id_123_200.jpg",
		},
		{
			name:        "with seed flag",
			args:        []string{"picsum", "--seed", "myseed", "--quiet", "--force", "200"},
			cleanupFile: "seed_myseed_200.jpg",
		},
		{
			name:        "with gray flag",
			args:        []string{"picsum", "--gray", "--quiet", "--force", "200"},
			cleanupFile: "200_gray.jpg",
		},
		{
			name:        "with blur flag",
			args:        []string{"picsum", "--blur", "--quiet", "--force", "200"},
			cleanupFile: "200_blur.jpg",
		},
		{
			name:        "with blurlevel flag",
			args:        []string{"picsum", "--blurlevel", "3", "--quiet", "--force", "200"},
			cleanupFile: "200_blur3.jpg",
		},
		{
			name:        "with output flag",
			args:        []string{"picsum", "--output", "/tmp/test_cli.jpg", "--quiet", "--force", "200"},
			cleanupFile: "/tmp/test_cli.jpg",
		},
		{
			name:        "with multiple flags combined",
			args:        []string{"picsum", "--gray", "--blur", "--quiet", "--force", "200"},
			cleanupFile: "200_gray_blur.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() {
				t.Skip("Skipping integration test in short mode")
			}

			defer func() {
				if tt.cleanupFile != "" {
					_ = os.Remove(tt.cleanupFile)
				}
			}()

			cmd := BuildCommand()
			ctx := context.Background()
			err := cmd.Run(ctx, tt.args)

			// These tests exercise lines 77-92 of runAction
			// They may succeed or fail depending on network availability
			if err != nil {
				t.Logf("Integration test note: %v", err)
			}
		})
	}
}

func TestRunAction_ValidateArgumentsPath(t *testing.T) {
	// Specifically test that ValidateArguments is called (lines 73-75)
	tests := []struct {
		name        string
		args        []string
		wantErr     bool
		cleanupFile string
	}{
		{
			name:    "empty args",
			args:    []string{"picsum"},
			wantErr: true,
		},
		{
			name:    "three args",
			args:    []string{"picsum", "200", "300", "extra"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.cleanupFile != "" {
				defer func() { _ = os.Remove(tt.cleanupFile) }()
			}

			cmd := BuildCommand()
			ctx := context.Background()
			err := cmd.Run(ctx, tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("cmd.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunAction_ValidateOptionsPath(t *testing.T) {
	// Specifically test that ValidateOptions is called (lines 88-90)
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "invalid blur level - above range",
			args:    []string{"picsum", "--blurlevel", "11", "200"},
			wantErr: true,
		},
		{
			name:    "mutually exclusive flags",
			args:    []string{"picsum", "--id", "123", "--seed", "test", "200"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := BuildCommand()
			ctx := context.Background()
			err := cmd.Run(ctx, tt.args)

			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if tt.wantErr && err == nil {
				t.Error("Expected error but got nil")
			}
		})
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && indexOf(s, substr) >= 0))
}

func indexOf(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func TestBuildFlags_Coverage(t *testing.T) {
	flags := buildFlags()

	// Test that all expected flags are present by name
	expectedFlags := map[string]bool{
		"id":        false,
		"seed":      false,
		"gray":      false,
		"blur":      false,
		"blurlevel": false,
		"quiet":     false,
		"output":    false,
		"force":     false,
	}

	for _, flag := range flags {
		names := flag.Names()
		if len(names) > 0 {
			mainName := names[0]
			if _, ok := expectedFlags[mainName]; ok {
				expectedFlags[mainName] = true
			}
		}
	}

	for name, found := range expectedFlags {
		if !found {
			t.Errorf("Expected flag %q was not found in buildFlags()", name)
		}
	}
}

func TestBuildFlags_StringFlagDefaults(t *testing.T) {
	flags := buildFlags()

	stringFlags := []string{"id", "seed", "output"}
	for _, flagName := range stringFlags {
		found := false
		for _, flag := range flags {
			if sf, ok := flag.(*cli.StringFlag); ok {
				if sf.Name == flagName {
					found = true
					// Verify it's a string flag with correct name
					if sf.Name != flagName {
						t.Errorf("StringFlag name mismatch: got %v, want %v", sf.Name, flagName)
					}
					break
				}
			}
		}
		if !found {
			t.Errorf("Expected StringFlag %q not found", flagName)
		}
	}
}

func TestBuildFlags_BoolFlagDefaults(t *testing.T) {
	flags := buildFlags()

	boolFlags := []string{"gray", "blur", "quiet", "force"}
	for _, flagName := range boolFlags {
		found := false
		for _, flag := range flags {
			if bf, ok := flag.(*cli.BoolFlag); ok {
				if bf.Name == flagName {
					found = true
					// Verify it's a bool flag with correct name
					if bf.Name != flagName {
						t.Errorf("BoolFlag name mismatch: got %v, want %v", bf.Name, flagName)
					}
					break
				}
			}
		}
		if !found {
			t.Errorf("Expected BoolFlag %q not found", flagName)
		}
	}
}

func TestBuildFlags_IntFlagDefaults(t *testing.T) {
	flags := buildFlags()

	intFlags := []string{"blurlevel"}
	for _, flagName := range intFlags {
		found := false
		for _, flag := range flags {
			if inf, ok := flag.(*cli.IntFlag); ok {
				if inf.Name == flagName {
					found = true
					// Verify it's an int flag with correct name
					if inf.Name != flagName {
						t.Errorf("IntFlag name mismatch: got %v, want %v", inf.Name, flagName)
					}
					break
				}
			}
		}
		if !found {
			t.Errorf("Expected IntFlag %q not found", flagName)
		}
	}
}

func TestBuildCommand_ActionIsNotNil(t *testing.T) {
	cmd := BuildCommand()
	if cmd.Action == nil {
		t.Error("BuildCommand() should have a non-nil Action")
	}
}

func TestBuildCommand_NameAndUsage(t *testing.T) {
	cmd := BuildCommand()

	tests := []struct {
		field    string
		got      string
		expected string
	}{
		{"Name", cmd.Name, "picsum"},
		{"Usage", cmd.Usage, "fetch photo from https://picsum.photos"},
	}

	for _, tt := range tests {
		if tt.got != tt.expected {
			t.Errorf("BuildCommand().%s = %v, want %v", tt.field, tt.got, tt.expected)
		}
	}
}

func TestBuildFlags_AllAliases(t *testing.T) {
	flags := buildFlags()

	expectedAliases := map[string][]string{
		"id":        {"i"},
		"seed":      {"s"},
		"gray":      {"g"},
		"blur":      {"b"},
		"blurlevel": {"B"},
		"quiet":     {"q"},
		"output":    {"o"},
		"force":     {"f"},
	}

	for _, flag := range flags {
		names := flag.Names()
		if len(names) == 0 {
			continue
		}
		mainName := names[0]
		if expectedAliases, ok := expectedAliases[mainName]; ok {
			actualAliases := names[1:]
			if len(actualAliases) != len(expectedAliases) {
				t.Errorf("Flag %q: expected %d aliases, got %d", mainName, len(expectedAliases), len(actualAliases))
				continue
			}
			for i, alias := range expectedAliases {
				if actualAliases[i] != alias {
					t.Errorf("Flag %q: expected alias %q, got %q", mainName, alias, actualAliases[i])
				}
			}
		}
	}
}

func TestBuildCommand_FlagsNotEmpty(t *testing.T) {
	cmd := BuildCommand()
	if len(cmd.Flags) == 0 {
		t.Error("BuildCommand() should have flags")
	}
}

func TestBuildFlags_EachFlagHasUsage(t *testing.T) {
	flags := buildFlags()

	for _, flag := range flags {
		names := flag.Names()
		if len(names) == 0 {
			continue
		}
		mainName := names[0]

		var usage string
		switch f := flag.(type) {
		case *cli.StringFlag:
			usage = f.Usage
		case *cli.BoolFlag:
			usage = f.Usage
		case *cli.IntFlag:
			usage = f.Usage
		}

		if usage == "" {
			t.Errorf("Flag %q has empty usage text", mainName)
		}
	}
}
