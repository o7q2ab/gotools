package main

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
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

	infs := rows(dirs)
	res := strings.Join(infs, "\n")
	fmt.Println(res)
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

func rows(dirs []string) []string {
	pids := []string{"PID"}
	govs := []string{"Go version"}
	deps := []string{"# Deps"}
	pths := []string{"Path"}
	mods := []string{"Module"}

	for i := range dirs {
		_, err := strconv.Atoi(dirs[i])
		if err != nil {
			continue
		}
		path, err := os.Readlink("/proc/" + dirs[i] + "/exe")
		if err != nil {
			continue
		}
		info, err := buildinfo.ReadFile(path)
		if err != nil {
			continue
		}

		pids = append(pids, dirs[i])
		govs = append(govs, info.GoVersion)
		deps = append(deps, strconv.Itoa(len(info.Deps)))
		pths = append(pths, path)
		mods = append(mods, fmt.Sprintf("%s %s", info.Main.Path, info.Main.Version))
	}

	if len(pids) == 0 {
		return []string{}
	}

	alignr(pids)
	alignr(govs)
	alignr(deps)
	alignl(pths)
	alignl(mods)

	infs := make([]string, len(pids))

	for i := range pids {
		infs[i] = fmt.Sprintf(
			"%s | %s | %s | %s | %s ",
			pids[i], govs[i], deps[i], pths[i], mods[i],
		)
	}

	return infs
}

func alignr(in []string) {
	longest := len(in[0])
	for i := 1; i < len(in); i++ {
		if len(in[i]) > longest {
			longest = len(in[i])
		}
	}

	diff := 0
	for i := range in {
		diff = longest - len(in[i])
		if diff == 0 {
			continue
		}
		for j := 0; j < diff; j++ {
			in[i] = " " + in[i]
		}
	}
}

func alignl(in []string) {
	longest := len(in[0])
	for i := 1; i < len(in); i++ {
		if len(in[i]) > longest {
			longest = len(in[i])
		}
	}

	diff := 0
	for i := range in {
		diff = longest - len(in[i])
		if diff == 0 {
			continue
		}
		for j := 0; j < diff; j++ {
			in[i] += " "
		}
	}
}
