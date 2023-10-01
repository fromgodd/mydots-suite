package main

import (
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    // Determine the user's home directory
    homeDir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println("Error determining home directory:", err)
        return
    }

    // Define the installation directory
    installDir := filepath.Join(homeDir, "mydots")

    // Create the installation directory if it doesn't exist
    if err := os.MkdirAll(installDir, os.ModePerm); err != nil {
        fmt.Println("Error creating installation directory:", err)
        return
    }

    // Copy files from the source directory to the installation directory
    if err := copyFiles(".", installDir); err != nil {
        fmt.Println("Error copying files:", err)
        return
    }

    fmt.Printf("mydots has been installed in %s\n", installDir)
}

// copyFiles copies files and directories from src to dst.
func copyFiles(src, dst string) error {
    return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        relPath, err := filepath.Rel(src, path)
        if err != nil {
            return err
        }

        destPath := filepath.Join(dst, relPath)

        if info.IsDir() {
            return os.MkdirAll(destPath, os.ModePerm)
        }

        if !info.Mode().IsRegular() {
            return nil // Skip non-regular files (e.g., directories, symlinks)
        }

        srcFile, err := os.Open(path)
        if err != nil {
            return err
        }
        defer srcFile.Close()

        destFile, err := os.Create(destPath)
        if err != nil {
            return err
        }
        defer destFile.Close()

        _, err = io.Copy(destFile, srcFile)
        if err != nil {
            return err
        }

        return nil
    })
}
