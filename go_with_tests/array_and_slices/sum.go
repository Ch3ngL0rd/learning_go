package array_and_slice

func Sum(numbers []int) int {
	sum := 0
	// Enumeration over range of numbers
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum [][]int) []int {
	return nil
}