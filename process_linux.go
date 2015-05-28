// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

func nameOf(pid int) string {
	procName := ""

	statusFile, err := procFsOpenPid(pid, "status")
	if err == nil {
		defer statusFile.Close()
		procName = procFsParseProcName(statusFile)
	}
	
	return procName
}

func count() int {
	return len(procFsListPids())
}

func listPids() []int {
	return procFsListPids()
}
