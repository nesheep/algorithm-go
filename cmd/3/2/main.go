package main

func calc(as []int, v int) int {
	cnt := 0
	for _, a := range as {
		if a == v {
			cnt++
		}
	}
	return cnt
}
