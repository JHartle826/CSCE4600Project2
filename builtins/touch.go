package builtins

import (
    "os"
)

func TouchFile(args ...string) error {
    // Create files
    for _, file := range args {
        _, err := os.Create(file)
        if err != nil {
            return err
        }
    }
    return nil
}
