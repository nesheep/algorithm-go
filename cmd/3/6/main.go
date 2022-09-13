package main

import "math"

func calc(k, n int) int {
	result := 0
	min := int(math.Min(float64(k), float64(n)))
	for x := 0; x <= min; x++ {
		for y := 0; y <= min; y++ {
			z := n - x - y
			if z >= 0 && z <= k {
				result++
			}
		}
	}
	return result
}
