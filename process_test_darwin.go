// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

func testCasesNameOf() []struct {
	in int
	want string} {

	cases := []struct {
		in int
		want string
	} {
		{38, "syslogd"},
		{1, "launchd"},
	}

	return cases
}

func testCasesPidOf() []struct {
	in string
	want int} {

	cases := []struct {
		in string
		want int
	} {
		{"syslogd", 38},
		{"launchd", 1},
	}

	return cases
}
