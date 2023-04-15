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
	procDir, err := os.Open("/proc")
	if err != nil {
		return fmt.Errorf("open '/proc' failed: %v", err)
	}
	defer procDir.Close()
	dirs, err := procDir.Readdirnames(-1)
	if err != nil {
		return err
	}
	skipped := 0
	readLinkFailed := 0
	notGo := 0
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	fmt.Fprintf(w, "[PID]\t| Go version\t| # Deps\t| Path\t| Mod\n")
	fmt.Fprintf(w, "-----\t| ----------\t| ------\t| ----\t| ---\n")
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

		fmt.Fprintf(w, "[%d]\t| %s\t| %d\t| %s\t| %s %s\n",
			pid, info.GoVersion, len(info.Deps), string(path), info.Main.Path, info.Main.Version)
	}
	w.Flush()
	fmt.Println("-------------")
	fmt.Printf("Total: %d, skipped: %d, read link failed: %d, not Go: %d\n", len(dirs), skipped, readLinkFailed, notGo)
	return nil
}
