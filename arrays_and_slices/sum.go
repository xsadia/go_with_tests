package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	Sum([]int{1, 2, 3, 4, 5})
}
