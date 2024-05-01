package builtins_test

import (
    "bytes"
    "github.com/jah0766/CSCE4600/Project2/builtins"
    "os"
    "testing"
)

func TestPrintWorkingDirectory(t *testing.T) {
    // Create a buffer to capture output
    var buf bytes.Buffer

    // Run the function
    err := builtins.PrintWorkingDirectory(&buf)

    // Check for error
    if err != nil {
        t.Errorf("PrintWorkingDirectory() returned an error: %v", err)
    }

    // Check the output captured in buf
    expected, _ := os.Getwd()
    if buf.String() != expected+"\n" {
        t.Errorf("PrintWorkingDirectory() produced unexpected output: got %q, want %q", buf.String(), expected)
    }
}
