package main

/*
	Smallest multiple
	2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

	What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

	https://projecteuler.net/problem=5
*/
func main() {

}

func SmallestMultiple(from, to int) int {
	i := to
	for {
		if Dividable(i, from, to) {
			return i
		}

		i += to
	}
}

func Dividable(num, from, to int) bool {
	for i := from; i <= to; i += 1 {
		if num%i != 0 {
			return false
		}
	}

	return true
}

func SmallestMultiple2(from, to int) int {
	res := 1
	for i := from; i <= to; i++ {
		res = lcm1(res, i)
	}

	return res
}

func lcm1(from int, to int) int {
	return (from * to) / GCD1(from, to)
}

func GCD1(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD1(b, a%b)
}

func SmallestMultiple3(from, to int) int {
	res := 1
	for i := from; i <= to; i++ {
		res = lcm2(res, i)
	}

	return res
}

func SmallestMultiple4(from, to int) int {
	res := 1
	for i := from; i <= to; i++ {
		res = lcm3(res, i)
	}

	return res
}

func lcm2(from int, to int) int {
	return (from * to) / GCD2(from, to)
}

func lcm3(from int, to int) int {
	return (from * to) / BinaryGCD(from, to)
}

func GCD2(a, b int) int {
	for {
		rem := a % b
		if rem == 0 {
			return b
		}

		a = b
		b = rem
	}
}

/*
http://www.blackwasp.co.uk/SteinsAlgorithm.aspx
*/
func BinaryGCD(a, b int) int {
	if a == b {
		return a
	}

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	if a%2 == 0 && b%2 == 0 {
		return BinaryGCD(a>>1, b>>1) << 1
	}

	if a%2 == 0 {
		return BinaryGCD(a>>1, b)
	}

	if b%2 == 0 {
		return BinaryGCD(a, b>>1)
	}

	if a >= b {
		return BinaryGCD((a-b)>>1, b)
	}

	return BinaryGCD(a, (b-a)>>1)
}
