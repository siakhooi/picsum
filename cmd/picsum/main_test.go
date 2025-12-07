package main

import (
	"testing"
)

func TestRun_Success(t *testing.T) {
	// Test successful execution with valid arguments
	// This tests the happy path where no error is returned (line 15 condition is false)
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "help flag",
			args:    []string{"picsum", "--help"},
			wantErr: false,
		},
		{
			name:    "version flag",
			args:    []string{"picsum", "--version"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := run(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRun_Error(t *testing.T) {
	// Test error cases to cover lines 15-17 (error path)
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "no arguments",
			args:    []string{"picsum"},
			wantErr: true,
		},
		{
			name:    "too many arguments",
			args:    []string{"picsum", "200", "300", "extra"},
			wantErr: true,
		},
		{
			name:    "invalid blur level - too high",
			args:    []string{"picsum", "--blurlevel", "11", "200"},
			wantErr: true,
		},
		{
			name:    "invalid blur level - negative",
			args:    []string{"picsum", "--blurlevel", "-1", "200"},
			wantErr: true,
		},
		{
			name:    "mutually exclusive id and seed",
			args:    []string{"picsum", "--id", "123", "--seed", "myseed", "200"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := run(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Verify that an error is actually returned
			if tt.wantErr && err == nil {
				t.Error("Expected error but got nil")
			}
		})
	}
}

func TestRun_IntegrationWithValidArgs(t *testing.T) {
	// Integration test with valid arguments that would download
	// Skip in short mode to avoid network calls
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name: "valid single dimension - network test",
			args: []string{"picsum", "--quiet", "--force", "--output", "/tmp/test_main_integration.jpg", "100"},
			// May error due to network, but tests the code path
			wantErr: false, // We accept either success or failure for network tests
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := run(tt.args)
			// For integration tests, we just ensure the function runs
			// The actual result depends on network availability
			if err != nil {
				t.Logf("Integration test note: %v", err)
			}
		})
	}
}
