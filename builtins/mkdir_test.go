package builtins_test

import (
    "os"
    "testing"
)

func TestMakeDirectory(t *testing.T) {
    // Define test directories
    testDirs := []string{"test_dir1", "test_dir2"}

    // Run the function
    err := builtins.MakeDirectory(testDirs...)

    // Check for error
    if err != nil {
        t.Errorf("MakeDirectory() returned an error: %v", err)
    }

    // Check if directories are created
    for _, dir := range testDirs {
        _, err := os.Stat(dir)
        if err != nil {
            t.Errorf("MakeDirectory() failed to create directory: %v", dir)
        }
        // Clean up: remove created directories
        defer os.Remove(dir)
    }
}
