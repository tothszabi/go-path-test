package main

import (
	"fmt"
	"github.com/bitrise-io/go-utils/pathutil"
	"os"
	"path/filepath"
)

func main() {
	paths := []string{
		"./",
		"./scripts/ci/bitrise/setup_bitrise_steps",
	}

	fmt.Printf("Running path expansion with Bitrise stuff\n")

	for _, path := range paths {
		expandedPath, err := pathutil.AbsPath(path)
		if err != nil {
			fmt.Printf("%s encoundtered an error: %s\n", path, err)
			continue
		}

		fmt.Printf("%s -> %s\n", path, expandedPath)
	}

	fmt.Printf("\nRunning path expansion with the GO runtime")

	for _, path := range paths {
		expandedPath, err := filepath.Abs(os.ExpandEnv(path))
		if err != nil {
			fmt.Printf("%s encoundtered an error: %s\n", path, err)
			continue
		}

		fmt.Printf("%s -> %s\n", path, expandedPath)
	}

	fmt.Printf("\nRunning path expansion with custom stuff\n")

	for _, path := range paths {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("current dir error: %s\n", err)
			continue
		}

		expandedPath := filepath.Join(currentDir, path)
		fmt.Printf("%s -> %s\n", path, expandedPath)
	}
}
