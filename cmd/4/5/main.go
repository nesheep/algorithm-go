package main

func count(n int) int {
	cnt := 0
	calc(n, 0, 0, &cnt)
	return cnt
}

func calc(n int, current int, use int, cnt *int) {
	if current > n {
		return
	}

	if use == 0b111 {
		*cnt++
	}

	calc(n, current*10+7, use|0b001, cnt)
	calc(n, current*10+5, use|0b010, cnt)
	calc(n, current*10+3, use|0b100, cnt)
}
