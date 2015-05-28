// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

// +build !darwin linux

package process

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const procFsPidPath = "/proc/%d/%s"

func openPid(pid int, name string) (*os.File, error) {
	return os.Open(fmt.Sprintf(procFsPidPath, pid, name))
}

func parseProcName(status *os.File) string {
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
