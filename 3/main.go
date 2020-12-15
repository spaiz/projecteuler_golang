package main

/*
	The prime factors of 13195 are 5, 7, 13 and 29.

	What is the largest prime factor of the number 600851475143 ?

	https://projecteuler.net/problem=3
 */
func main() {
	
}

func LPF1(num int) int {
	max := 0
	for i := 1; i*i < num; i++ {
		if num % i == 0 {
			if isPrime(i) {
				if i > max {
					max = i
				}
			}
		}
	}

	return max
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