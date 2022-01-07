package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, numberSlice := range numbers {
		sums = append(sums, Sum(numberSlice))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, numbersSlice := range numbers {
		if len(numbersSlice) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbersSlice[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}

func main() {
	Sum([]int{1, 2, 3, 4, 5})
}

/*
	x := [3]string{"Лайка", "Белка", "Стрелка"}

	y := x[:] // slice "y" points to the underlying array "x"

	z := make([]string, len(x))
	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

	y[1] = "Belka" // the value at index 1 is now "Belka" for both "y" and "x"
*/
