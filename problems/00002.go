package problems

import (
	"fmt"
)

type problem00002 struct {
	sum int
}

func (prob problem00002) getName() string {
	return "Even Fibonacci Numbers"
}

func (prob *problem00002) solve() {
	prob.sum = 0

	prev := 1
	curr := 2

	for curr < 4000000 {
		if curr%2 == 0 {
			prob.sum += curr
		}

		next := prev + curr
		prev = curr
		curr = next
	}
}

func (prob problem00002) getResult() string {
	return fmt.Sprintf("%v", prob.sum)
}
