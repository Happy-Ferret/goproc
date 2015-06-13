// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

// +build !darwin linux

package procfs

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

const (
	root = "/proc"
	path = "/proc/%s"
	pidPath = "/proc/%d/%s"
)

type Stat struct {
	Name  string
	UTime int
	STime int
	VSize int
}

const (
	statName   = 1
	statUTime  = 13
	statSTime  = 14
	statVmSize = 22
	statHighestIndex = statVmSize
)

func newStat() *Stat {
	stat := new(Stat)
	stat.Name = ""
	stat.UTime = -1
	stat.STime = -1
	return stat
}


func ListPids() []int {
	items, err := ioutil.ReadDir(root)
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

func CpuTimeTotal() int {
	stat, err := os.Open(fmt.Sprintf(path, "stat"))
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

func StatOf(pid int) *Stat {
	result := newStat()

	stat, err := os.Open(fmt.Sprintf(pidPath, pid, "stat"))
	defer stat.Close()
	if err != nil {
		return result
	}
	scanner := bufio.NewScanner(stat)
	if !scanner.Scan() {
		return result
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < statHighestIndex {
		return result
	}
	result.Name = parts[statName][1 : len(parts[statName])-1] // strip '(', ')'
	result.UTime = atoiOr(parts[statUTime], -1)
	result.STime = atoiOr(parts[statSTime], -1)
	result.VSize = atoiOr(parts[statVmSize], -1)

	return result
}

func atoiOr(s string, alt int) int {
	value, err := strconv.Atoi(s)
	if err == nil {
		return value
	}
	return alt
}
