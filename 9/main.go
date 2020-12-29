package main

/*
	Special Pythagorean triplet

	a^2 + b^2 = c^2
	For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

	There exists exactly one Pythagorean triplet for which a + b + c = 1000.
	Find the product abc.

	https://projecteuler.net/problem=9
*/
func main() {

}

type Result struct {
	a int
	b int
	c int
}

func (r Result) Product() int {
	return r.a * r.b * r.c
}

func SpecialPythagoreanTriplet1(sum int) *Result {
	res := &Result{
		a: 0,
		b: 0,
		c: 0,
	}

	for a := 1; a <= sum; a++ {
		for b := a + 1; b <= sum-a; b++ {
			for c := b + 1; c <= sum-a-b; c++ {
				if a*a+b*b == c*c {
					if  a + b + c == sum {
						res.a = a
						res.b = b
						res.c = c
						return res
					}
				}
			}
		}
	}

	return res
}

func SpecialPythagoreanTriplet2(sum int) *Result {
	res := &Result{
		a: 0,
		b: 0,
		c: 0,
	}

	for a := 1; a <= sum; a++ {
		for b := a + 1; b <= sum-a; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				res.a = a
				res.b = b
				res.c = c
				return res
			}
		}
	}

	return res
}
