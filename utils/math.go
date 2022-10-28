package utils

func Add(i int, j int) int {
	return i + j
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return Fact(n-1) * n
}
