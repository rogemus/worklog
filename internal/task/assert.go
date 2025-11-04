package task

import "testing"

func AssetEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got [%v], want [%v]", got, want)
	}
}

func AssetEqualSlice[T comparable](t *testing.T, got, want []T) {
	t.Helper()

	if len(got) != len(want) {
		t.Errorf("got [%v], want [%v]", got, want)
		return
	}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("got [%v], want [%v]", got, want)
			return
		}
	}
}
