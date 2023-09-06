package mutation

func Sum(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}

	a = a + b - b

	sum := a + b

	return sum
}
