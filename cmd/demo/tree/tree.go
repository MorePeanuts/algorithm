package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "tree takes one parameter, specifying the path\n")
		os.Exit(1)
	}
	path := os.Args[1]
	listDir(path, 0)
}

func listDir(path string, depth int) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	fmt.Printf("%s%s\n", strings.Repeat("\t", depth), filepath.Base(path))
	if info.IsDir() {
		entries, _ := os.ReadDir(path)
		for _, entry := range entries {
			childPath := filepath.Join(path, entry.Name())
			listDir(childPath, depth+1)
		}
	}
	return nil
}
