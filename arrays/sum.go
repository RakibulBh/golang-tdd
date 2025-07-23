package arrays

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Finds the sum of every array passed in
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// Sums all the numbers in arrays except the first value
func SumAllTails(numbersToSumTails ...[]int) []int {
	sums := make([]int, len(numbersToSumTails))
	for i, arr := range numbersToSumTails {
		if len(arr) > 1 {
			sum := Sum(arr[1:])
			sums[i] = sum
		} else {
			sums[i] = 0
		}

	}

	return sums
}
