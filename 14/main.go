package main

/*
	Longest Collatz sequence

	The following iterative sequence is defined for the set of positive integers:

	n → n/2 (n is even)
	n → 3n + 1 (n is odd)

	Using the rule above and starting with 13, we generate the following sequence:

	13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
	It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

	Which starting number, under one million, produces the longest chain?

	NOTE: Once the chain starts the terms are allowed to go above one million.

	https://projecteuler.net/problem=14
*/
func main() {

}

type Result struct {
	chain []int
	start int
}

func (r *Result) Chain() []int {
	return r.chain
}

func (r *Result) ChainSize() int {
	return len(r.chain)
}

func (r *Result) GetStart() int {
	return r.start
}

func LongestCollatzSequence(start int) *Result {
	maxChain := &Result{}
	for i := 0; i <= start; i++ {
		res := CollatzSequence(i)
		if res.ChainSize() > maxChain.ChainSize() {
			maxChain = res
		}
	}

	return maxChain
}

func CollatzSequence(i int) *Result {
	res := &Result{
		start:i,
	}
	cur := i
	for cur > 1 {
		res.chain = append(res.chain, cur)
		if cur % 2 == 0 {
			cur /= 2
			continue
		}

		cur = 3*cur + 1
	}

	res.chain = append(res.chain, cur)

	return res
}


func LongestCollatzSequenceMemoize(start int) *Result {
	memo := map[int]*Result{}
	maxChain := &Result{}
	for i := 3; i <= start; i++ {
		res := CollatzSequenceMemoize(i, memo)
		for i, v := range res.chain {
			if memo[v] == nil {
				memo[v] = &Result{
					start:v,
				}
				memo[v].chain = append(memo[v].chain, res.chain[i:]...)
			}
		}
		if res.ChainSize() > maxChain.ChainSize() {
			maxChain = res
		}
	}

	return maxChain
}

func CollatzSequenceMemoize(i int, memo map[int]*Result) *Result {
	res := &Result{
		start:i,
	}

	cur := i
	for cur > 1 {
		if memo[cur] != nil {
			res.chain = append(res.chain, memo[cur].chain...)
			return res
		}

		res.chain = append(res.chain, cur)
		if cur % 2 == 0 {
			cur /= 2
			continue
		}

		cur = 3*cur + 1
	}

	res.chain = append(res.chain, cur)
	return res
}