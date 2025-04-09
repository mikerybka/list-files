package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "."
	if len(os.Args) >= 2 {
		root = os.Args[1]
	}

	// Walk the directory
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
