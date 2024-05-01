package builtins

import (
    "fmt"
    "io"
    "strings"
)

func Echo(w io.Writer, args ...string) error {
    // Join arguments with space
    msg := strings.Join(args, " ")
    
    // Print the message
    _, err := fmt.Fprintln(w, msg)
    return err
}
