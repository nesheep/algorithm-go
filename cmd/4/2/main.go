package main

import "fmt"

func main() {
	fmt.Println(tribonacci(10))
	fmt.Println(tribonacci(20))
	fmt.Println(tribonacci(30))
	fmt.Println(tribonacci(40))
	fmt.Println(tribonacci(50))
	fmt.Println(tribonacci(60))
}

func tribonacci(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}
	return calc(n, memo)
}

func calc(n int, memo []int) int {
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if memo[n] != -1 {
		return memo[n]
	}
	memo[n] = calc(n-1, memo) + calc(n-2, memo) + calc(n-3, memo)
	return memo[n]
}
