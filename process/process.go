// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

type Property int

const (
	VmUsage Property = iota
	CpuUsage
)

type PropertyMap map[Property]interface{}

func NameOf(pid int) string {
	return nameOf(pid)
}

func Count() int {
	return count()
}

func ListPids() []int {
	return listPids()
}

func PidOf(name string) int {
	result := -1

	for _, pid := range ListPids() {
		nameOfPid := NameOf(pid)
		if name == nameOfPid {
			result = pid
			break
		}
	}

	return result
}

func PropertiesOf(pid int, keys []Property) PropertyMap {
	return propertiesOf(pid, keys)
}
