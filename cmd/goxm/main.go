package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
)

const header = `%s

  - Package: %s
  - Module:  %s %s (%s)
  - Version: %s
  - Size:    %d bytes
  - Mode:    %s
`

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
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	info, err := buildinfo.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Printf(
		header, path, info.Path, info.Main.Path, info.Main.Version, info.Main.Sum,
		info.GoVersion, stat.Size(), stat.Mode(),
	)

	fmt.Printf("\n  - Build settings:\n")
	for _, s := range info.Settings {
		fmt.Printf("    %s=%s\n", s.Key, s.Value)
	}

	fmt.Printf("\n  - Dependencies [%d]:\n", len(info.Deps))
	for _, d := range info.Deps {
		fmt.Printf("    %s %s\n", d.Path, d.Version)
	}
	return nil
}
