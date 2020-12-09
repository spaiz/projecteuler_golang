package main

import "fmt"

/*
	Multiples of 3 and 5

	If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
	Find the sum of all the multiples of 3 or 5 below 1000.

	https://projecteuler.net/problem=1
*/

func main() {
	n := 1000
	res := FindIter(n)
	fmt.Printf("Result: %d\n", res)

	res2 := FindRec(n - 1)
	fmt.Printf("Result: %d\n", res2)
}

func FindIter(num int) int {
	var sum int
	for i := 1; i < num; i++ {
		rem1 := i % 3
		if rem1 == 0 {
			sum += i
			continue
		}

		rem2 := i % 5
		if rem2 == 0 {
			sum += i
		}
	}

	return sum
}

func FindRec(num int) int {
	if num == 0 {
		return 0
	}

	rem1 := num % 3
	if rem1 == 0 {
		return num + FindRec(num - 1)
	}

	rem2 := num % 5
	if rem2 == 0 {
		return num + FindRec(num - 1)
	}

	return 0 + FindRec(num - 1)
}