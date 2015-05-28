// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

/*
#include <stdlib.h>
#include "libproc.h"
 
int get_proc_name(int pid, char* name, int name_size) {
  return proc_name(pid, name, name_size);
}

int get_proc_count() {
  return proc_listallpids(NULL, 0);
}
*/
import "C"
import "unsafe"
import "strings"

func nameOf(pid int) string {
	name := C.CString(strings.Repeat("\x00", 1024))
	defer C.free(unsafe.Pointer(name))
	nameLen := C.get_proc_name(C.int(pid), name, C.int(1024))
	var result string

	if nameLen > 0 {
		result = C.GoString(name);
	} else {
		result = ""
	}

	return result
}

func count() int {
	procs := int(C.get_proc_count())
	if procs <= 0 {
		return 0
	} else {
		return procs
	}
}

func listPids() []int {
	return make([]int, 0)
}