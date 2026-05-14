package tests

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSampleFilesExist(t *testing.T) {
	paths := []string{
		"../samples/basic/hello/main.go",
		"../samples/basic/variables/main.go",
		"../samples/basic/return_value/main.go",
		"../samples/unsupported/control_flow/main.go",
		"../samples/unsupported/functions/main.go",
	}

	for _, path := range paths {
		if _, err := os.Stat(filepath.Clean(path)); err != nil {
			t.Fatalf("expected sample file %s to exist: %v", path, err)
		}
	}
}
