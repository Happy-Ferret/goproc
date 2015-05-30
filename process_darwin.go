// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

/*
#include <stdlib.h>
#include "libproc.h"
*/
import "C"
import "unsafe"
import "strings"
//import "strconv"
//import "fmt"

func nameOf(pid int) string {
	name := C.CString(strings.Repeat("\x00", 1024))
	namePtr := unsafe.Pointer(name)
	defer C.free(namePtr)
	nameLen := C.proc_name(C.int(pid), namePtr, C.uint32_t(1024))
	var result string

	if nameLen > 0 {
		result = C.GoString(name)
	} else {
		result = ""
	}

	return result
}

func count() int {
	procs := int(C.proc_listallpids(nil, C.int(0)))
	if procs <= 0 {
		return 0
	} else {
		return procs
	}
}

func listPids() []int {
	cnt := count()
	if cnt <= 0 {
		return make([]int, 0)
	}
	pids := C.malloc(C.size_t(cnt * int(unsafe.Sizeof(C.int(0)))))
	if pids == nil {
		return make([]int, 0)
	}
	defer C.free(unsafe.Pointer(pids))
	pidsCnt := C.proc_listallpids(pids, C.int(cnt*int(unsafe.Sizeof(C.int(0)))))
	if pidsCnt <= 0 {
		return make([]int, 0)
	}
	casted := (*[1 << 20]C.int)(unsafe.Pointer(pids))
	pidsCopy := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		pidsCopy[i] = int(casted[i])
	}
	return trimPidArray(pidsCopy)
}

// garbage collectable type of interesting C.struct_proc_taskinfo fields
type processInfo struct {
	pid int
	virtualSize int64
}

func propertiesOf(pid int, keys []int) PropertyMap {
	result := make(PropertyMap)
	var thread processInfoHandler = threadInfoHandler
	var task processInfoHandler = taskInfoHandler
	chain := thread.compose(task) 
	taskInfo := processInfoOf(pid, chain)
	for _, key := range keys {
		switch key {
		case VMUsage:
			result[VMUsage] = taskInfo.virtualSize
		}
	}
	return result
}

// as in https://gist.github.com/gotohr/7005197
type processInfoHandler func (info *processInfo) *processInfo

func (f processInfoHandler) compose(inner processInfoHandler) processInfoHandler {
	return func(info *processInfo) *processInfo { return f(inner(info)) }
}

func processInfoOf(pid int, handler processInfoHandler) *processInfo {
	info := new(processInfo)
	info.pid = pid
	handler(info)
	return handler(info)
}

func threadInfoHandler(info *processInfo) *processInfo {
	// just a stub, TODO: complete
	return info
}

func taskInfoHandler(info *processInfo) *processInfo {
	taskInfo := C.malloc(C.size_t(C.PROC_PIDTASKINFO_SIZE))
	defer C.free(taskInfo)
	size := C.proc_pidinfo(C.int(info.pid), C.PROC_PIDTASKINFO, 0, taskInfo, C.int(C.PROC_PIDTASKINFO_SIZE))
	
	// checking size as described in http://goo.gl/Lta0IO
	if size < C.int(unsafe.Sizeof(taskInfo)) {
		return info
	}
	
	casted := (*C.struct_proc_taskinfo)(taskInfo)
	info.virtualSize = int64(casted.pti_virtual_size) // bytes

	return info
}

func trimPidArray(pids []int) []int {
	index := len(pids)
	for i, e := range pids {
		if e <= 0 {
			index = i
			break
		}
	}
	return pids[0:index]
}
