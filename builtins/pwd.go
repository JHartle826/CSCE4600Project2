package builtins

import (
    "fmt"
    "io"
    "os"
)

func PrintWorkingDirectory(w io.Writer) error {
    // Get current working directory
    wd, err := os.Getwd()
    if err != nil {
        return err
    }
    
    // Print the current working directory
    _, err = fmt.Fprintln(w, wd)
    return err
}
