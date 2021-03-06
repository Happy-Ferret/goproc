// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

func testCasesNameOf() []struct {
	in   int
	want string
} {

	cases := []struct {
		in   int
		want string
	}{
		{1, "init"},
		{562, "sshd"},
	}

	return cases
}

func testCasesPidOf() []struct {
	in   string
	want int
} {

	cases := []struct {
		in   string
		want int
	}{
		{"init", 1},
		{"sshd", 562},
	}

	return cases
}
