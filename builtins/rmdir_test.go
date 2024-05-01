package builtins_test

import (
    "os"
    "testing"
)

func TestRemoveDirectory(t *testing.T) {
    // Define test directories
    testDir := "test_dir"

    // Create test directory
    err := os.Mkdir(testDir, 0755)
    if err != nil {
        t.Fatalf("Failed to create test directory: %v", err)
    }

    // Run the function
    err = builtins.RemoveDirectory(testDir)

    // Check for error
    if err != nil {
        t.Errorf("RemoveDirectory() returned an error: %v", err)
    }

    // Check if directory is removed
    _, err = os.Stat(testDir)
    if !os.IsNotExist(err) {
        t.Errorf("RemoveDirectory() failed to remove directory: %v", testDir)
    }
}
