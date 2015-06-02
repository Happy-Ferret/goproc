// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

// +build !darwin linux

package process

import (
	"fmt"
	"testing"
)

func TestProcFsCpuTimeTotal(t *testing.T) {
	got := procFsCpuTimeTotal()
	if got < 0 {
		t.Errorf("procFsCpuTimeTotal() > 0 == false, want true")
	} else {
		fmt.Printf("procFsCpuTimeTotal() == %v\n", got)
	}
}

/*
func TestProcFsJiffiesOf(t *testing.T) {
	utimeGot, stimeGot := procFsJiffiesOf(1)
	if utimeGot < 0 || stimeGot < 0 {
		t.Errorf("utime,stime => procFsJiffiesOf() > 0 == false, want true")
	} else {
		fmt.Printf("utime,stime => procFsJiffiesOfl() == %v, %v\n", utimeGot, stimeGot)
	}
*/
