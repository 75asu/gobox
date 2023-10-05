package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    // Define command-line flags
    directory := flag.String("dir", ".", "Directory path for file renaming")
    rules := flag.String("rules", ".txt:_renamed,.jpg:_renamed", "File renaming rules (comma-separated)")

    flag.Parse()

    renamingRules := parseRules(*rules)

    err := filepath.Walk(*directory, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Check if the file is not a directory
        if !info.IsDir() {
            // Get the file extension
            ext := filepath.Ext(path)

            // Check if there is a renaming rule for the file extension
            if newExt, ok := renamingRules[ext]; ok {
                // Define the new file name
                newName := strings.TrimSuffix(path, ext) + newExt + ext

                // Rename the file
                err := os.Rename(path, newName)
                if err != nil {
                    fmt.Printf("Error renaming %s: %v\n", path, err)
                } else {
                    fmt.Printf("Renamed %s to %s\n", path, newName)
                }
            }
        }

        return nil
    })

    if err != nil {
        fmt.Println("Error:", err)
    }
}

func parseRules(rulesStr string) map[string]string {
    rules := make(map[string]string)
    rulePairs := strings.Split(rulesStr, ",")

    for _, rulePair := range rulePairs {
        parts := strings.Split(rulePair, ":")
        if len(parts) == 2 {
            rules[parts[0]] = parts[1]
        }
    }

    return rules
}
