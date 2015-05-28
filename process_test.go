// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

import "testing"

func TestNameOf(t *testing.T) {

	cases := testCasesNameOf()
	
	for _, c := range cases {
		got := NameOf(c.in)
		if got != c.want {
			t.Errorf("NameOf(%d) == %s, want %s", c.in, got, c.want)
		}
	}
}
