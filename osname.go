package main

import (
	"fmt"
	"runtime"
	"bytes"
	"os"
	"os/exec"
	"strings"
	"time"
)

type GoInfoObject struct {
	GoOS 		string
	Kernel 		string
	Core 		string
	Platform 	string
	OS 			string
	Hostname 	string
	CPUs 		int
}

func (gi *GoInfoObject) VarDump() {
	fmt.Println("GoOS:", gi.GoOS)
	fmt.Println("Kernel:", gi.Kernel)
	fmt.Println("Core:", gi.Core)
	fmt.Println("Platform:", gi.Platform)
	fmt.Println("OS:", gi.OS)
	fmt.Println("Hostname:", gi.Hostname)
	fmt.Println("CPUs:", gi.CPUs)
}

func GetInfo() (GoInfoObject, error) {
	kernel, core, platform, osname, err := _getInfo("-srio")
	gio := GoInfoObject{
		Kernel:   kernel,
		Core:     core,
		Platform: platform,
		OS:       osname,
		GoOS:     runtime.GOOS,
		CPUs:     runtime.NumCPU(),
		Hostname: _hostname(),
	}
	return gio, err
}

func _getInfo(flags string) (string, string, string, string, error) {
	out, err := _uname(flags)
	tries := 0
	for strings.Contains(out, "broken pipe") && tries < 3 {
		out, err = _uname(flags)
		time.Sleep(500 * time.Millisecond)
		tries++
	}
	if strings.Contains(out, "broken pipe") {
		out = ""
	}
	kernel, core, platform, osname := _expandInfo(out)
	return kernel, core, platform, osname, err
}

func _uname(flags string) (string, error) {
	cmd := exec.Command("uname", flags)
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	return out.String(), err
}

func _expandInfo(osStr string) (string, string, string, string) {
	osStr = strings.Replace(osStr, "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	osInfo := strings.Split(osStr, " ")
	kernel := "Unknown"
	core := "Unknown"
	platform := "Unknown"
	osname := "Unknown"
	if len(osInfo) > 0 {
		kernel = osInfo[0]
	}
	if len(osInfo) > 1 {
		core = osInfo[1]
	}
	if len(osInfo) > 2 {
		platform = osInfo[2]
	}
	if len(osInfo) > 3 {
		osname = osInfo[3]
	}
	return kernel, core, platform, osname
}

func _hostname() string {
	host, _ := os.Hostname()
	return host
}