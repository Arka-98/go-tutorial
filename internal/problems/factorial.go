package problems

func Factorial(input int) int {
	if input == 0 {
		return 1
	}

	return input * Factorial(input-1)
}
