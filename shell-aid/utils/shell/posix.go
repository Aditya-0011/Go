//go:build !windows

package shell

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	objects "shell-aid/lib"
	"strconv"
	"strings"
)

func psNameAndPPID(pid int) (name string, ppid int, err error) {
	out, err := exec.Command("ps", "-o", "comm=", "-o", "ppid=", "-p", strconv.Itoa(pid)).Output()
	if err != nil {
		return "", 0, err
	}
	fields := strings.Fields(strings.TrimSpace(string(out)))
	if len(fields) < 2 {
		return "", 0, exec.ErrNotFound
	}
	ppid, _ = strconv.Atoi(fields[len(fields)-1])
	name = strings.Join(fields[:len(fields)-1], " ")
	return name, ppid, nil
}

func detectShell() string {
	known := map[string]string{
		"bash": "bash", "zsh": "zsh", "fish": "fish",
		"sh": "sh", "dash": "dash", "ksh": "ksh",
		"tcsh": "tcsh", "csh": "csh",
		"nu": "nushell", "xonsh": "xonsh",
	}

	pid := os.Getpid()
	for i := 0; i < 20 && pid > 0; i++ {
		name, ppid, err := psNameAndPPID(pid)
		if err != nil {
			break
		}
		n := strings.ToLower(strings.TrimLeft(filepath.Base(name), "-"))
		if s, ok := known[n]; ok {
			return s
		}
		pid = ppid
	}

	if sh := os.Getenv("SHELL"); sh != "" {
		return filepath.Base(sh)
	}
	return "unknown"
}

func Detect() objects.SystemInfo {
	return objects.SystemInfo{
		OS:    runtime.GOOS,
		Shell: detectShell(),
	}
}
