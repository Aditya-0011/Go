//go:build windows

package shell

import (
	"os"
	objects "shell-aid/lib"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

type procInfo struct {
	pid  int
	ppid int
	exe  string
}

func snapshotProcs() (map[int]procInfo, error) {
	h, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(h)

	var e windows.ProcessEntry32
	e.Size = uint32(unsafe.Sizeof(e))

	procs := make(map[int]procInfo, 256)

	if err = windows.Process32First(h, &e); err != nil {
		return nil, err
	}
	for {
		name := strings.ToLower(windows.UTF16ToString(e.ExeFile[:]))
		procs[int(e.ProcessID)] = procInfo{
			pid:  int(e.ProcessID),
			ppid: int(e.ParentProcessID),
			exe:  name,
		}
		if err = windows.Process32Next(h, &e); err != nil {
			if err == windows.ERROR_NO_MORE_FILES {
				break
			}
			return nil, err
		}
	}
	return procs, nil
}

func detectShell() string {
	procs, err := snapshotProcs()
	if err == nil {
		pid := os.Getpid()
		for range 20 {
			p, ok := procs[pid]
			if !ok {
				break
			}
			switch p.exe {
			case "cmd.exe", "cmd":
				return "cmd"
			case "powershell.exe", "powershell", "pwsh.exe", "pwsh":
				return "powershell"
			case "nu.exe", "nu":
				return "nushell"
			}
			if p.ppid == 0 || p.ppid == pid {
				break
			}
			pid = p.ppid
		}
	}

	if strings.Contains(strings.ToLower(os.Getenv("COMSPEC")), "cmd.exe") {
		return "cmd"
	}
	if os.Getenv("POWERSHELL_DISTRIBUTION_CHANNEL") != "" || os.Getenv("PSModulePath") != "" {
		return "powershell"
	}
	return "unknown"
}

func Detect() objects.SystemInfo {
	return objects.SystemInfo{
		OS:    "windows",
		Shell: detectShell(),
	}
}
