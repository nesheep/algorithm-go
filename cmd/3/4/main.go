package main

const inf = 20000000

func calc(as []int) int {
	max := -inf
	min := inf
	for _, a := range as {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	return max - min
}
