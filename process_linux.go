// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

import "strings"
import "fmt"

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
      items := procFsParseStatusItems(pid, []string{"VmSize"})
      if len(items) == 1 {
        vmUsage := AtoiOr(strings.Fields(strings.TrimSpace(items[0]))[0], -1) // bytes
        if vmUsage > 0 {
          result[VmUsage] = vmUsage * 1000 // bytes
          break
        }
      }
			result[VmUsage] = -1

		case CpuUsage:
			result[CpuUsage] = -1000
		}
	}

	return result
}
