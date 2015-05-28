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
