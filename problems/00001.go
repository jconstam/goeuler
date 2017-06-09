package problems

import (
	"fmt"
)

type problem00001 struct {
	sum int
}

func (prob problem00001) getName() string {
	return "Multiples of 3 and 5"
}

func (prob *problem00001) solve() {
	prob.sum = 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			prob.sum += i
		}
	}
}

func (prob problem00001) getResult() string {
	return fmt.Sprintf("%v", prob.sum)
}
