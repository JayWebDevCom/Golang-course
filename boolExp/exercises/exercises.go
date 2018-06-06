package exercises

// Exercise1 is exported  - half and isEXEN
func Exercise1(num int) (float64, bool) {
	return float64(num) / 2, float64(num)/2 == 0
}

// Exercise2 is exported - half and isEXEN
func Exercise2(num int, myCall func(int) bool) (float64, bool) {
	return float64(num) / 2, myCall(num)
}

// Exercise3 is exported - greatest in variadic
func Exercise3(num ...float64) float64 {
	var ret float64
	for _, e := range num {
		if e > ret {
			ret = e
		}
	}
	return ret
}

// Exercise4 is exported - length of variadic
func Exercise4(num ...float64) int {
	return len(num)
}

// project euler
