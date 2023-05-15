package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	paths := []string{
		"",
		"",
	}

	fmt.Println("Running path expansion with the GO runtime")

	for _, path := range paths {
		expandedPath, err := filepath.Abs(os.ExpandEnv(path))
		if err != nil {
			fmt.Printf("%s encoundtered an error: %s", path, err)
			continue
		}

		fmt.Printf("%s -> %s", path, expandedPath)
	}

	fmt.Println("")
	fmt.Println("Running path expansion with custom stuff")

	for _, path := range paths {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("current dir error: %s", err)
			continue
		}

		expandedPath := filepath.Join(currentDir, path)
		fmt.Printf("%s -> %s", path, expandedPath)
	}
}
