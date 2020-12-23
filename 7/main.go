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

	if num%2 == 0 {
		return false
	}

	for i := 3; i*i <= num; i = i + 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func PrimeNst(n int) int {
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

func PrimeSeieve1(n int) int {
	const arrSize = 1000000
	numbers := [arrSize]int{}

	j := 2
	for i := 0; i < len(numbers); i++ {
		numbers[i] = j
		j++
	}

	for i, p := range numbers {
		if p == -1 {
			continue
		}

		for j := i; j < len(numbers); j += p {
			if numbers[j] != p {
				numbers[j] = -1
			}
		}
	}

	found := 0
	for _, p := range numbers {
		if p != -1 {
			found++
			if found == n {
				return p
			}
		}
	}

	return -1
}

/*
	Using indexes as a numbers
 */
func PrimeSeieve2(n int) int {
	const max = 1000000
	numbers := [max]bool{}

	for i := 2; i*i < len(numbers); i++ {
		for y := i*i; y < len(numbers); y += i {
			numbers[y] = true
		}
	}

	nthPrime := 0

	for i := 2; i < len(numbers); i++ {
		if !numbers[i] {
			nthPrime++
			if nthPrime == n {
				return i
			}
		}
	}

	return -1
}

/**
	Using bitmask as optimization for memory
 */
func PrimeSeieve3(n int) int {
	const max = 1000000
	numbers := NewBitSet(max)

	for i := 2; i*i < numbers.Size(); i++ {
		for y := i*i; y < numbers.Size(); y += i {
			numbers.Set(y, true)
		}
	}

	nthPrime := 0

	for i := 2; i < numbers.Size(); i++ {
		if !numbers.Get(i) {
			nthPrime++
			if nthPrime == n {
				return i
			}
		}
	}

	return -1
}

type BitSet []uint8

func NewBitSet(n int) BitSet {
	return make(BitSet, (n+7)/8)
}

func (b BitSet) Get(index int) bool {
	pos := index / 8
	j := index % 8
	return (b[pos] & (uint8(1) << j)) != 0
}

func (b BitSet) Set(index int, value bool) {
	pos := index / 8
	j := uint(index % 8)
	if value {
		b[pos] |= uint8(1) << j
	} else {
		b[pos] &= ^(uint8(1) << j)
	}
}

func (b BitSet) Size() int {
	return 8 * len(b)
}