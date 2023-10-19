package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbers2Sum ...[]int) []int {
	var sums []int

	for _, numbers := range numbers2Sum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {

			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
