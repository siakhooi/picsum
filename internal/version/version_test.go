package version

import "testing"

func TestGetVersion(t *testing.T) {
	expected := "0.0.1"
	actual := GetVersion()

	if actual != expected {
		t.Errorf("GetVersion() = %q, want %q", actual, expected)
	}
}
