package fibonacci

func FibonacciFor(n int) int {
	a, b := 0, 1

	if n == a {
		return a
	}
	for i := 2; i <= n; i++ {
		c := a + b
		a, b = b, c
	}
	return b
}

func FinonacciRec(n int) int {
	if n == 1 {
		return n
	}
	return FinonacciRec(n-1) + FinonacciRec(n-2)
}

var fibonacciData = []int{0, 1}

func fibonacciRecMemo(n int) int {
	if n == 0 {
		return fibonacciData[0]
	}
	if len(fibonacciData) >= n+1 {
		return fibonacciData[1]
	}
	data := fibonacciRecMemo(n-1) + fibonacciRecMemo(n-2)

	fibonacciData = append(fibonacciData, data)

	return data
}
