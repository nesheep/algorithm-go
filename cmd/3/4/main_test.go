package main

import "testing"

func TestCalc(t *testing.T) {
	as := []int{33, 42, 36, 68}

	got := calc(as)
	want := 35

	if got != want {
		t.Errorf("expect %d, but %d", want, got)
	}
}
