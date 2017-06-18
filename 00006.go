package goeuler

type problem00006 struct {
	difference float64
}

func (prob problem00006) getInfo() (string, string) {
	return "Sum square difference", "The sum of the squares of the first ten natural numbers is,\n1^2 + 2^2 + ... + 10^2 = 385\nThe square of the sum of the first ten natural numbers is,\n(1 + 2 + ... + 10)^2 = 55^2 = 3025\nHence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum."
}

func (prob *problem00006) solve() {
	sum := 0
	squareSum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		squareSum += i * i
	}

	prob.difference = float64((sum * sum) - squareSum)
}

func (prob problem00006) getExpectedResult() float64 {
	return float64(25164150)
}

func (prob problem00006) getActualResult() float64 {
	return prob.difference
}
