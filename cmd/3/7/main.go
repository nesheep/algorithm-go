package main

func calc(s string) int {
	n := len(s)
	result := 0
	for bit := 0; bit < (1 << (n - 1)); bit++ {
		tmp := 0
		for i := 0; i < n-1; i++ {
			tmp *= 10
			tmp += int(s[i] - '0')
			if bit&(1<<i) != 0 {
				result += tmp
				tmp = 0
			}
		}
		tmp *= 10
		tmp += int(s[len(s)-1] - '0')
		result += tmp
	}
	return result
}
