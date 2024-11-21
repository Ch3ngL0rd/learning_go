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
		if len(numbers) == 0 {
			tail_sums = append(tail_sums, 0)
		} else {
			tail_sums = append(tail_sums, Sum(numbers[1:]))
		}
	}
	return tail_sums
}