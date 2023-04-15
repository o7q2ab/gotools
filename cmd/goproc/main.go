package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"text/tabwriter"
)

func main() {
	if runtime.GOOS != "linux" {
		fmt.Println("Linux only.")
		os.Exit(1)
	}

	if err := run(); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	dirs, err := procDirs()
	if err != nil {
		return err
	}

	skipped, readLinkFailed, notGo := 0, 0, 0

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	row := func(format string, a ...any) { fmt.Fprintf(w, format, a...) }

	row("PID\t| Go version\t| # Deps\t| Path\t| Module\n")
	row("---\t| ----------\t| ------\t| ----\t| ------\n")

	for i := range dirs {
		pid, err := strconv.Atoi(dirs[i])
		if err != nil {
			skipped++
			continue
		}
		path, err := os.Readlink("/proc/" + dirs[i] + "/exe")
		if err != nil {
			readLinkFailed++
			continue
		}
		info, err := buildinfo.ReadFile(path)
		if err != nil {
			notGo++
			continue
		}

		row("%d\t| %s\t| %d\t| %s\t| %s %s\n",
			pid, info.GoVersion, len(info.Deps), string(path), info.Main.Path, info.Main.Version)
	}
	w.Flush()
	fmt.Println("-------------")
	fmt.Printf("Total: %d, skipped: %d, read link failed: %d, not Go: %d\n", len(dirs), skipped, readLinkFailed, notGo)
	return nil
}

func procDirs() ([]string, error) {
	d, err := os.Open("/proc")
	if err != nil {
		return nil, err
	}
	defer d.Close()
	return d.Readdirnames(-1)
}
