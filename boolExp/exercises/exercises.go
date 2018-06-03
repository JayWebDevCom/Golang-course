package exercises

// Exercise1 is exported
func Exercise1(num int) (float64, bool) {
	return float64(num) / 2, float64(num)/2 == 0
}
