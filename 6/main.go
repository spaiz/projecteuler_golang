package main

/*
	Sum square difference
	The sum of the squares of the first ten natural numbers is,
		1^2 + 2^2 + ... + 10^2 = 385
	The square of the sum of the first ten natural numbers is,
		(1 + 2 + ... + 10)^2 = 55^2 = 3025
	Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is
		3025 - 385 = 2640.

	Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

https://projecteuler.net/problem=6
*/
func main() {
	
}

func Diff1(n int) int {
	sumOfSquares := 0
	sum := 0
	for i := 1; i <= n; i++ {
		sumOfSquares += i*i
		sum += i
	}

	return sum*sum - sumOfSquares
}

func Diff2(n int) int {
	sumOfSquares := (n*(n+1)*(2*n+1))/6
	sum := n * (n+1) / 2
	return sum*sum - sumOfSquares
}