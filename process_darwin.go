// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

/*
#include <stdlib.h>
#include "libproc.h"
 
int call_proc_name(int pid, char* name, int name_size) {
  return proc_name(pid, name, name_size);
}
*/
import "C"
import "unsafe"
import "strings"

func nameOf(pid int) string {
	name := C.CString(strings.Repeat("\x00", 1024))
	defer C.free(unsafe.Pointer(name))
	nameLen := C.call_proc_name(C.int(pid), name, C.int(1024))
	var result string

	if nameLen > 0 {
		result = C.GoString(name);
	} else {
		result = ""
	}

	return result
}
