package array_and_slice

func Sum(numbers []int) int {
	sum := 0
	// Enumeration over range of numbers
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var tail_sums []int
	for _, numbers := range numbersToSum {
		sum := Sum(numbers)
		if len(numbers) > 1 {
			sum -= numbers[0]
		}
		tail_sums = append(tail_sums, sum)
	}
	return tail_sums
}