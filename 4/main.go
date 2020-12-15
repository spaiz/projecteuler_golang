package main

/*
	A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

	Find the largest palindrome made from the product of two 3-digit numbers.

	https://projecteuler.net/problem=4
*/
func main() {

}

func LargestPalindrome1(n int) (int, int, int) {
	maxNum := 1
	for i := 0; i < n; i++ {
		maxNum *= 10
	}

	maxNum -= 1
	lastI := 0
	lastY := 0
	maxPal := 0

	for i := 1; i <= maxNum; i++ {
		for y := 1; y <= maxNum; y++ {
			candidate := i * y
			if isPalindrome(candidate) {
				if candidate > maxPal {
					lastI = i
					lastY = y
					maxPal = candidate
				}
			}
		}
	}

	return maxPal, lastI, lastY
}

func isPalindrome(num int) bool {
	reversed := 0
	origin := num
	for num > 0 {
		digit := num % (10)
		num /= 10
		reversed = reversed*10 + digit
	}

	return reversed == origin
}
