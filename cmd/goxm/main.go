package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: '%s <path>'\n", os.Args[0])
		os.Exit(1)
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run(path string) error {
	info, err := buildinfo.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(path)
	fmt.Printf("  %s\n", info.GoVersion)
	fmt.Printf("  %s\n", info.Path)
	fmt.Printf("  Dependencies:\n")
	for _, d := range info.Deps {
		fmt.Printf("    %s (%s)\n", d.Path, d.Version)
	}
	fmt.Printf("  Build settings:\n")
	for _, s := range info.Settings {
		fmt.Printf("    %s=%s\n", s.Key, s.Value)
	}
	return nil
}
