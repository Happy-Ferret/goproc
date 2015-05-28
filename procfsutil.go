// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

// +build !darwin linux

package process

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"strconv"
)

const procFsRoot = "/proc"
const procFsPidPath = "/proc/%d/%s"

func procFsOpenPid(pid int, name string) (*os.File, error) {
	return os.Open(fmt.Sprintf(procFsPidPath, pid, name))
}

func procFsParseProcName(status *os.File) string {
	procName := ""

	scanner := bufio.NewScanner(status)
	if scanner.Scan() {
		if parts := strings.Split(scanner.Text(), ":"); len(parts) == 2 {
			procName = strings.TrimSpace(parts[1])
		}
	}
	//if err := scanner.Err(); err != nil {
	//}

	return procName
}

func procFsListPids() []int {
	items, err := ioutil.ReadDir(procFsRoot)
	if err != nil {
		return []int{}
	}
	
	pids := make([]int, len(items))
	pids[0] = -1 // mark value
	i := 0
	for _,item := range items {
		pid := procFsTryNameToPid(item.Name())
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

func procFsTryNameToPid(name string) int {
	pid, err := strconv.Atoi(name)
	if err != nil  || pid <= 0 {
		return -1
	}
	
	return pid
}
