# goproc
Process management API for Go.

## Project Status
This is a stub. No more than _explorative programming_ for the moment.
Work in progress...

## Example
```go
package main

import "github.com/gsscoder/goproc/process"

func main() {
  processName := process.NameOf(1) // result: "init" (on Linux)
  processId := process.PidOf("launchd") // result: 1 (on OS X)
  count := process.Count() // result: int count of running processes
  pids := process.ListPids() // result: []int array with running pids
  props := process.PropertiesOf(PidOf("Xorg"), []Property{VmUsage, CpuUsage}) // result: map[Property]interface{}
  xorgVm := props[VmUsage] // int (bytes)
  xorgCpu := props[CpuUsage] // float32 (%), this value for now is correct only under Linux
}
```

## Tests
Depending on function (``process.NameOf()`` for example) and platform type you may need run as root.
```sh
cd /path/to/goproc/process
sudo go test
```
