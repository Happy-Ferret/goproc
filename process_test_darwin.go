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
