// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

// +build !darwin linux

package process

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	procFsRoot = "/proc"
	procFsPidPath = "/proc/%d/%s"
	procFsPath = "/proc/%s"
)

type procFsStat struct {
	name  string
	utime int
	stime int
	vsize int
}

const (
	procFsStatName   = 1
	procFsStatUTime  = 13
	procFsStatSTime  = 14
	procFsStatVmSize = 22
	procFsStatHighestIndex = procFsStatVmSize 
)

func newProcFsStat() *procFsStat {
	stat := new(procFsStat)
	stat.name = ""
	stat.utime = -1
	stat.stime = -1
	return stat
}

func procFsOpenPid(pid int, name string) (*os.File, error) {
	return os.Open(fmt.Sprintf(procFsPidPath, pid, name))
}

func procFsOpen(name string) (*os.File, error) {
	return os.Open(fmt.Sprintf(procFsPath, name))
}

func procFsParseStatusItems(pid int, keys []string) []string {
	status, err := procFsOpenPid(pid, "status")
	if err != nil {
		return make([]string, 0)
	}
	defer status.Close()

	values := make([]string, len(keys))
	i := 0

	scanner := bufio.NewScanner(status)
	for scanner.Scan() {
		if parts := strings.Split(scanner.Text(), ":"); len(parts) == 2 {
			if currkey := strings.TrimSpace(parts[0]); strElemIndexOf(currkey, keys) >= 0 {
				values[i] = strings.TrimSpace(parts[1])
				i++
			}
		}
	}

	return values
}

func procFsListPids() []int {
	items, err := ioutil.ReadDir(procFsRoot)
	if err != nil {
		return []int{}
	}

	pids := make([]int, len(items))
	pids[0] = -1 // mark value
	i := 0
	for _, item := range items {
		pid := atoiOr(item.Name(), -1)
		if pid > 0 {
			pids[i] = pid
			i++
		}
	}

	if pids[0] > 0 { // some pid added
		return pids[0:i]
	}
	return []int{}
}

func procFsCpuTimeTotal() int {
	stat, err := procFsOpen("stat")
	if err != nil {
		return -1
	}
	defer stat.Close()
	scanner := bufio.NewScanner(stat)
	if !scanner.Scan() {
		return -1
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return -1
	}
	if strings.TrimSpace(parts[0]) != "cpu" {
		return -1
	}
	total := 0
	for _, cpuTime := range parts[1:] {
		partial := atoiOr(cpuTime, -1)
		if partial < 0 {
			return -1
		}
		total += partial
	}
	return total
}

func procFsStatOf(pid int) *procFsStat {
	result := newProcFsStat()

	stat, err := procFsOpenPid(pid, "stat")
	defer stat.Close()
	if err != nil {
		return result
	}
	scanner := bufio.NewScanner(stat)
	if !scanner.Scan() {
		return result
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < procFsStatHighestIndex {
		return result
	}
	result.name = parts[procFsStatName][1 : len(parts[procFsStatName])-1] // strip '(', ')'
	result.utime = atoiOr(parts[procFsStatUTime], -1)
	result.utime = atoiOr(parts[procFsStatSTime], -1)
	result.vsize = atoiOr(parts[procFsStatVmSize], -1)

	return result
}
