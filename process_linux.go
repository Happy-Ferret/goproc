// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

/*
#include <unistd.h>
*/
import "C"
import "time"
//import "strings"
//import "fmt"

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
			result[CpuUsage] = cpuUsageOf(pid, func() {time.Sleep(time.Second)})
		}
	}

	return result
}

func cpuCount() int {
	return int(C.sysconf(C._SC_NPROCESSORS_ONLN))
}


func cpuUsageOf(pid int, waitHandler func()) float32 {
	// as explained in http://goo.gl/fjrV16
	stat1 := procFsStatOf(pid)
	utime1 := stat1.utime
	stime1 := stat1.stime
	cputime1 := procFsCpuTimeTotal()

	waitHandler()

	stat2 := procFsStatOf(pid)
	utime2 := stat2.utime
	stime2 := stat2.stime
	cputime2 := procFsCpuTimeTotal()

	return float32(cpuCount() * ((utime2 + stime2) - (utime1 + stime1)) * 100) / float32(cputime2 - cputime1)
}
