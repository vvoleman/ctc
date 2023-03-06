package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	cases := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	f := Fibonacci()

	for i, want := range cases {
		result := f()

		if result != want {
			t.Errorf("(run #%d) Fibonnaci() == %d, expects %d", i, result, want)
		}
	}
}
