package builtins_test

import (
    "bytes"
    "github.com/jah0766/CSCE4600Project2/builtins"
    "testing"
)

func TestEcho(t *testing.T) {
    // Create a buffer to capture output
    var buf bytes.Buffer

    // Run the function
    err := builtins.Echo(&buf, "Hello", "World")

    // Check for error
    if err != nil {
        t.Errorf("Echo() returned an error: %v", err)
    }

    // Check the output captured in buf
    expected := "Hello World\n"
    if buf.String() != expected {
        t.Errorf("Echo() produced unexpected output: got %q, want %q", buf.String(), expected)
    }
}
