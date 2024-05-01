package builtins

import (
    "errors"
    "os"
)

func RemoveDirectory(args ...string) error {
    // Remove directories
    for _, dir := range args {
        err := os.Remove(dir)
        if err != nil {
            return errors.New("failed to remove directory: " + dir)
        }
    }
    return nil
}
