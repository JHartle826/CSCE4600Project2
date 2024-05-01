package builtins_test

import (
    "os"
    "testing"
)

func TestTouchFile(t *testing.T) {
    // Define test files
    testFiles := []string{"test_file1.txt", "test_file2.txt"}

    // Run the function
    err := builtins.TouchFile(testFiles...)

    // Check for error
    if err != nil {
        t.Errorf("TouchFile() returned an error: %v", err)
    }

    // Check if files are created
    for _, file := range testFiles {
        _, err := os.Stat(file)
        if err != nil {
            t.Errorf("TouchFile() failed to create file: %v", file)
        }
        // Clean up: remove created files
        defer os.Remove(file)
    }
}
