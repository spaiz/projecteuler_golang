package main

/*
	10001st prime

	By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

	What is the 10 001st prime number?

	https://projecteuler.net/problem=7
*/
func main() {
	
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	if num == 2 {
		return true
	}

	if num % 2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i = i + 2 {
		if num % i == 0 {
			return false
		}
	}

	return true
}

func Prime_Nst(n int) int {
	found := 0
	num := 0
	i := 1
	for found < n {
		if isPrime(i) {
			found++
			num = i
		}

		i++
	}

	return num
}