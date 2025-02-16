package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
Package procmeminfo provides an interface for /proc/meminfo

	import "github.com/guillermo/go.procmeminfo"

	meminfo := &procmeminfo.MemInfo{}
	meminfo.Update()

Once the info was updated you can access like a normal map[string]float64

	v := (*meminfo)["MemTotal"]  // 1809379328 (1766972 * 1024)

It also implement some handy methods, like:

	meminfo.Total() // (*meminfo)["MemTotal"]
	meminfo.Free()  // MemFree + Buffers + Cached
	meminfo.Used()  // Total - Free


Return all the values in units, so while you get this from cat /proc/meminfo

	MemTotal:        1766972 kB
	MemFree:          115752 kB
	Buffers:            3172 kB
	Cached:           182552 kB
	SwapCached:        83572 kB
	Active:          1055284 kB
	Inactive:         382872 kB
	Active(anon):     932712 kB
	Inactive(anon):   329508 kB
	Active(file):     122572 kB
	Inactive(file):    53364 kB
	Unevictable:       10640 kB
	Mlocked:           10640 kB
	SwapTotal:       1808668 kB
	SwapFree:        1205672 kB
	Dirty:               100 kB
	Writeback:             0 kB
	AnonPages:       1214740 kB
	Mapped:           115636 kB
	Shmem:              4840 kB
	Slab:              77412 kB
	SReclaimable:      34344 kB
	SUnreclaim:        43068 kB
	KernelStack:        4328 kB
	PageTables:        39428 kB
	NFS_Unstable:          0 kB
	Bounce:                0 kB
	WritebackTmp:          0 kB
	CommitLimit:     2692152 kB
	Committed_AS:    5448372 kB
	VmallocTotal:   34359738367 kB
	VmallocUsed:      106636 kB
	VmallocChunk:   34359618556 kB
	HardwareCorrupted:     0 kB
	AnonHugePages:         0 kB
	HugePages_Total:       0
	HugePages_Free:        0
	HugePages_Rsvd:        0
	HugePages_Surp:        0
	Hugepagesize:       2048 kB
	DirectMap4k:      216236 kB
	DirectMap2M:     1593344 kB


All the kB values are multiply by 1024
*/

type GoMemInfo struct {
	MemTotal		uint64
	MemAvailable 	uint64
	MemUsed 		uint64
	MemSwap 		uint64
}

type MemInfo map[string]uint64
// Update s with current values, usign the pid stored in the Stat
func (m *MemInfo) Update() error {
	var err error

	path := filepath.Join("/proc/meminfo")
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		n := strings.Index(text, ":")
		if n == -1 {
			continue
		}

		key := text[:n] // metric
		data := strings.Split(strings.Trim(text[(n+1):], " "), " ")
		if len(data) == 1 {
			value, err := strconv.ParseUint(data[0], 10, 64)
			if err != nil {
				continue
			}
			(*m)[key] = value
		} else if len(data) == 2 {
			if data[1] == "kB" {
				value, err := strconv.ParseUint(data[0], 10, 64)
				if err != nil {
					continue
				}
				(*m)[key] = value * 1024
			}
		}

	}
	return nil
}

func ConvertToStruct(memInfo *MemInfo) GoMemInfo {
	if memInfo == nil {
		return GoMemInfo{}
	}

	d := *memInfo
	memFree := d["MemFree"]
	buffers := d["Buffers"]
	cached := d["Cached"]
	memTotal := d["MemTotal"]
	swapTotal := d["SwapTotal"]
	swapFree := d["SwapFree"]

	// Calculate available memory
	memAvailable := memFree + buffers + cached

	// Calculate used memory
	memUsed := memTotal - memAvailable

	// Calculate swap usage percentage, avoiding division by zero
	var swapUsage uint64
	if swapTotal > 0 {
		swapUsage = uint64((100 * (swapTotal - swapFree)) / swapTotal)
	}

	return GoMemInfo{
		MemTotal:     memTotal,
		MemAvailable: memAvailable,
		MemUsed:      memUsed,
		MemSwap:      swapUsage,
	}
}
