package main

import "testing"

func TestCalc(t *testing.T) {
	as := []int{1, 2, 3, 1, 1, 1, 2, 2, 3, 4, 0}
	v := 1

	got := calc(as, v)
	want := 4

	if got != want {
		t.Errorf("expect %d, but %d", want, got)
	}
}
