package version

import "testing"

const expectedPicsumVersion = "0.3.0"

func TestGetVersion(t *testing.T) {
	actual := GetVersion()
	expected := expectedPicsumVersion

	if actual != expected {
		t.Errorf("GetVersion() = %q, want %q", actual, expected)
	}
}
