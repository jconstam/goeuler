package goeuler

type problem00001 struct {
	sum float64
}

func (prob problem00001) getInfo() (string, string) {
	return "Multiples of 3 and 5", "If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.\nFind the sum of all the multiples of 3 or 5 below 1000."
}

func (prob *problem00001) solve() {
	prob.sum = 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			prob.sum += float64(i)
		}
	}
}

func (prob problem00001) getExpectedResult() float64 {
	return float64(233168)
}

func (prob problem00001) getActualResult() float64 {
	return prob.sum
}
