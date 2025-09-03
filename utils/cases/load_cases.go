package cases

import (
	"os"
	"testing"
)

func LoadCases(t *testing.T, path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read %s: %v", path, err)
	}

	return data
}
