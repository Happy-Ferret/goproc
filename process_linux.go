// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

//import "strings"
//import "fmt"

//func nameOf(pid int) string {
//	items := procFsParseStatusItems(pid, []string{"Name"})
//	if len(items) != 1 {
//		return ""
//	}
//	return items[0]
//}

func nameOf(pid int) string {
	return procFsStatOf(pid).name
}


func count() int {
	return len(procFsListPids())
}

func listPids() []int {
	return procFsListPids()
}

func propertiesOf(pid int, keys []int) PropertyMap {
	result := make(PropertyMap)
	stat := procFsStatOf(pid)

	for _, key := range keys {
		switch key {
		case VmUsage:
			result[VmUsage] = stat.vsize

		case CpuUsage:
			result[CpuUsage] = -1000
		}
	}

	return result
}
