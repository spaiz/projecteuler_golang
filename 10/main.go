package main

/*
	Summation of primes
	The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

	Find the sum of all the primes below two million.

	https://projecteuler.net/problem=10
*/
func main() {

}

/*
	Using Seieve algorithm with bit mask, for memory efficiency.
 */
func SumOfPrimesBelow1(upperLimit int) int {
	numbers := NewBitSet(upperLimit)

	for i := 2; i*i < numbers.Size(); i++ {
		for y := i*i; y < numbers.Size(); y += i {
			numbers.Set(y, true)
		}
	}

	sum := 0

	for i := 2; i < upperLimit; i++ {
		if !numbers.Get(i) {
			sum += i
		}
	}

	return sum
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