package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"runtime"
	"strconv"
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

		fmt.Printf("[%d]\t%s | deps: %d\t| %s\n", pid, info.GoVersion, len(info.Deps), string(path))
	}
	fmt.Println("---")
	fmt.Printf("Total: %d, skipped: %d, read link failed: %d, not Go: %d\n", len(dirs), skipped, readLinkFailed, notGo)
	return nil
}
