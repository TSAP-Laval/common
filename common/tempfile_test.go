package common

import "testing"

type tempTestCase struct {
	in   bool
	want bool
}

func TestHello(t *testing.T) {
	cases := []tempTestCase{
		tempTestCase{true, false},
		tempTestCase{false, true},
	}

	for _, c := range cases {
		got := Hello(c.in)

		if got != c.want {
			t.Errorf("Hello(%t) == %t, want %t", c.in, got, c.want)
		}
	}
}
