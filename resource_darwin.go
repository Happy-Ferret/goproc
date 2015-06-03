package process

/*
#include <stdlib.h>
#include <mach/message.h>
#include <mach/mach_host.h>
#include <mach/host_info.h>
*/
import "C"
//import "unsafe"

func cpuTimeTotal() int {
	selfHost := C.mach_host_self()
	hostInfo := C.malloc(C.size_t(C.HOST_CPU_LOAD_INFO_COUNT))
	count := C.mach_msg_type_number_t(C.HOST_CPU_LOAD_INFO_COUNT)

	err := C.host_statistics(C.host_t(selfHost), C.HOST_CPU_LOAD_INFO, C.host_info_t(hostInfo), &count)
	defer C.free(hostInfo)

	if err != C.kern_return_t(C.KERN_SUCCESS) {
		return 0
	}

	return -1
}
