package main

const inf = 20000000

func calc(as []int) int {
	min := inf
	second := inf
	for _, a := range as {
		if a < min {
			second = min
			min = a
		} else if a < second {
			second = a
		}
	}
	return second
}
