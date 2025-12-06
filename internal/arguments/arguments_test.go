package arguments

import (
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
