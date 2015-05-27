package process

func testCasesNameOf() []struct {
	in int
	want string} {

	cases := []struct {
		in int
		want string
	} {
		{1, "init"},
		{562, "sshd"},
	}

	return cases
}
