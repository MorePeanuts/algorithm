package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "du takes one parameter, specifying the path\n")
		os.Exit(1)
	}
	path := os.Args[1]
	sizeDirectory(path)
}

func sizeDirectory(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "path %v does not exist", path)
		os.Exit(1)
	}
	totalSize := info.Size()
	if info.IsDir() {
		entries, _ := os.ReadDir(path)
		for _, entry := range entries {
			childPath := filepath.Join(path, entry.Name())
			totalSize += sizeDirectory(childPath)
		}
	}
	fmt.Printf("%-10d %s\n", totalSize, path)
	return totalSize
}
