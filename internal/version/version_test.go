package version

import "testing"

const expectedPicsumVersion = "1.0.0"

func TestGetVersion(t *testing.T) {
	actual := Version()
	expected := expectedPicsumVersion

	if actual != expected {
		t.Errorf("GetVersion() = %q, want %q", actual, expected)
	}
}
