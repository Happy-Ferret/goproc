// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

func nameOf(pid int) string {
	procName := ""

	statusFile, err := openPid(pid, "status")
	if err == nil {
		defer statusFile.Close()
		procName = parseProcName(statusFile)
	}
	
	return procName
}

func count() int {
	return len(listAllPids())
}

func listPids() []int {
	return listAllPids()
}
