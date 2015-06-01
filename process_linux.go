// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

func nameOf(pid int) string {
	procName := ""

	statusFile, err := procFsOpenPid(pid, "status")
	if err == nil {
		defer statusFile.Close()
		procName = procFsParseStatusItems(statusFile, []string{"Name"})[0]
	}

	return procName
}

func count() int {
	return len(procFsListPids())
}

func listPids() []int {
	return procFsListPids()
}

func propertiesOf(pid int, keys []int) PropertyMap {
	//panic("propertiesOf() for Linux not implemented")
	return make(PropertyMap)
}
