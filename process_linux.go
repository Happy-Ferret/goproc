// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

func nameOf(pid int) string {
 	items := procFsParseStatusItems(pid, []string{"Name"})
	if len(items) != 1 {
		return ""
	}
	return items[0]
}

func count() int {
	return len(procFsListPids())
}

func listPids() []int {
	return procFsListPids()
}

func propertiesOf(pid int, keys []int) PropertyMap {
	result := make(PropertyMap)

	for _,key := range keys {
		switch key {
		case VmUsage:
			result[VmUsage] = -1000
		case CpuUsage:
			result[CpuUsage] = -1000
		}
	}

	return result
}
