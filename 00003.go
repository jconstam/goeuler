package goeuler

type problem00003 struct {
	largestFactor float64
}

func (prob problem00003) getInfo() (string, string) {
	return "Largest prime largestFactor", "The prime factors of 13195 are 5, 7, 13 and 29.\nWhat is the largest prime factor of the number 600851475143 ?"
}

func (prob *problem00003) solve() {
	value := float64(600851475143)
	newValue := value

	counter := float64(2)
	for counter*counter <= newValue {
		if int64(newValue)%int64(counter) == 0 {
			newValue = newValue / counter
			prob.largestFactor = counter
		} else {
			counter++
		}
	}

	if newValue > prob.largestFactor {
		prob.largestFactor = newValue
	}
}

func (prob problem00003) getExpectedResult() float64 {
	return float64(6857)
}

func (prob problem00003) getActualResult() float64 {
	return prob.largestFactor
}
