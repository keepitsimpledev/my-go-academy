package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(slicesToSum ...[]int) []int {
	sums := []int{}
	for _, slice := range slicesToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(slicesToSum ...[]int) []int {
	sums := []int{}

	for _, slice := range slicesToSum {
		sum := 0

		if len(slice) > 0 {
			slice = slice[1:]
			sum = Sum(slice)
		}

		sums = append(sums, sum)
	}

	return sums
}
