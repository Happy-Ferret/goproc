// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

/*
#include <stdlib.h>
#include "libproc.h"
*/
import "C"
import "unsafe"
import "strings"
import "strconv"
//import "fmt"

type go_proc_taskinfo struct { //sys/proc_info.h/proc_taskinfo
	pti_virtual_size C.uint64_t			// virtual memory size (bytes)
	pti_resident_size C.uint64_t			// resident memory size (bytes)
	pti_total_user C.uint64_t				// total time
	pti_total_system C.uint64_t		
	pti_threads_user C.uint64_t			// existing threads only
	pti_threads_system C.uint64_t		
	pti_policy C.int32_t					// default policy for new threads
	pti_faults C.int32_t					// number of page faults
	pti_pageins C.int32_t					// number of actual pageins
	pti_cow_faults C.int32_t					// number of copy-on-write faults
	pti_messages_sent C.int32_t				// number of messages sent
	pti_messages_received C.int32_t				// number of messages received
	pti_syscalls_mach C.int32_t				// number of mach system calls
	pti_syscalls_unix C.int32_t				// number of unix system calls
	pti_csw C.int32_t			          	// number of context switches
	pti_threadnum C.int32_t					// number of threads in the task
	pti_numrunning C.int32_t					// number of running threads
	pti_priority C.int32_t					// task priority
}

func nameOf(pid int) string {
	name := C.CString(strings.Repeat("\x00", 1024))
	namePtr := unsafe.Pointer(name)
	defer C.free(namePtr)
	nameLen := C.proc_name(C.int(pid), namePtr, C.uint32_t(1024))
	var result string

	if nameLen > 0 {
		result = C.GoString(name);
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
	pidsCnt := C.proc_listallpids(pids, C.int(cnt * int(unsafe.Sizeof(C.int(0)))))
	if (pidsCnt <= 0) {
		return make([]int, 0)
	}
	casted := (*[1<<20]C.int)(unsafe.Pointer(pids))
	pidsCopy := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		pidsCopy[i] = int(casted[i])
	}
	return trimPidArray(pidsCopy)
}

func propertiesOf(pid int, keys []int) PropertyMap {
	result := make(PropertyMap)
	
	info := C.malloc(C.size_t(C.PROC_PIDTASKINFO_SIZE))
	defer C.free(info)
	actualSize := C.proc_pidinfo(C.int(pid), C.PROC_PIDTASKINFO, 0, info, C.int(C.PROC_PIDTASKINFO_SIZE)) //C.int(unsafe.Sizeof(info)))
	// checking size as described in http://goo.gl/Lta0IO
	if actualSize < C.int(unsafe.Sizeof(info)) {
		//panic(fmt.Sprintf("actualsize=%d\n, sizeof(info)=%d", int(actualSize), int(C.int(unsafe.Sizeof(info))))) //DEBUG
		return result
	}
	casted := (*go_proc_taskinfo)(info)
	
	for _,key := range keys {
		switch key {
		case PropertyVMSize:
			result[PropertyVMSize] = strconv.FormatInt(int64(casted.pti_virtual_size), 10)
		}
	}

	return result
}

func trimPidArray(pids []int) []int {
	index := len(pids)
	for i,e := range pids {
		if e <= 0 {
			index = i
			break
		}
	}
	return pids[0:index]
}