// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

/*
#include <unistd.h>
*/
import "C"
import "github.com/gsscoder/goproc/process/internal/procfs"
import "time"
//import "strings"
//import "fmt"

func nameOf(pid int) string {
	return procfs.StatOf(pid).Name
}


func count() int {
	return len(procfs.ListPids())
}

func listPids() []int {
	return procfs.ListPids()
}

func propertiesOf(pid int, keys []Property) PropertyMap {
	result := make(PropertyMap)
	stat := procfs.StatOf(pid)

	for _, key := range keys {
		switch key {
		case VmUsage:
			result[VmUsage] = stat.VSize

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
	stat1 := procfs.StatOf(pid)
	utime1 := stat1.UTime
	stime1 := stat1.STime
	cputime1 := procfs.CpuTimeTotal()

	waitHandler()

	stat2 := procfs.StatOf(pid)
	utime2 := stat2.UTime
	stime2 := stat2.STime
	cputime2 := procfs.CpuTimeTotal()

	return float32(cpuCount() * ((utime2 + stime2) - (utime1 + stime1)) * 100) / float32(cputime2 - cputime1)
}
