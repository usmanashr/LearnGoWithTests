package arrays

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(a []int, b []int) []int {
	result := make([]int, len(a))
	result[0] = Sum(a)
	result[1] = Sum(b)

	return result
}
