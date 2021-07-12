package calc

func Add(a, b int) int {
	return a + b
}

func AddMany(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}
