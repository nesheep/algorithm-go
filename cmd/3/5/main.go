package main

const inf = 20000000

func howManyTimes(a int) int {
	exp := 0
	for a%2 == 0 {
		a /= 2
		exp++
	}
	return exp
}

func calc(as []int) int {
	result := inf
	for _, a := range as {
		exp := howManyTimes(a)
		if exp < result {
			result = exp
		}
	}
	return result
}
