package problems

func SumOfDigits(n int) int {
	if n < 10 {
		return n
	}

	return n % 10 + SumOfDigits(n / 10)
}