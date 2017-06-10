package problems

type problem00005 struct {
	smallestMultiple float64
}

func (prob problem00005) getInfo() (string, string) {
	return "Smallest multiple", "2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.\nWhat is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?"
}

func (prob *problem00005) solve() {
	currVal := 2520
	for {
		currVal += 20
		good := false
		for i := 19; i > 1; i-- {
			if currVal%i == 0 {
				good = true
			} else {
				good = false
				break
			}
		}

		if good {
			prob.smallestMultiple = float64(currVal)
			break
		}
	}
}

func (prob problem00005) getExpectedResult() float64 {
	return float64(232792560)
}

func (prob problem00005) getActualResult() float64 {
	return prob.smallestMultiple
}
