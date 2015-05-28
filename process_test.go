// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

import (
	//"fmt"
	"testing"
)

func TestNameOf(t *testing.T) {

	cases := testCasesNameOf()
	
	for _, c := range cases {
		got := NameOf(c.in)
		if got != c.want {
			t.Errorf("NameOf(%d) == %s, want %s", c.in, got, c.want)
		}
	}
}

func TestCount(t *testing.T) {
	want := true
	got := Count();
	if got > 0 != want {
		t.Errorf("Count() > 0 == %t, want %t", got, want)
	}
	//else {
	//	fmt.Printf("Count() == %d\n", got)
	//}
}
