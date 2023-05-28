package firsttest

func Sum(a, b int) int {
	return a + b
}

func Cum(numbers ...int) int {
	resp := 0

	for _, v := range numbers {
		resp += v
	}
	return resp
}
