package problems

import "fmt"

type problem00003 struct {
	factor int64
}

func (prob problem00003) getName() string {
	return "Largest prime factor"
}

func (prob *problem00003) solve() {
	value := int64(600851475143)
	newValue := value

	counter := int64(2)
	for counter*counter <= newValue {
		if newValue%counter == 0 {
			newValue = newValue / counter
			prob.factor = counter
		} else {
			counter++
		}
	}

	if newValue > prob.factor {
		prob.factor = newValue
	}
}

func (prob problem00003) getResult() string {
	return fmt.Sprintf("%v", prob.factor)
}
