package goeuler

import (
	"testing"
)

// TestSolveAll
func TestSolveAll(t *testing.T) {
	main()
}

func TestAll(t *testing.T) {
	for i, currProblem := range Problems {
		currProblem.solve()

		actualResult := currProblem.getActualResult()
		expectedResult := currProblem.getExpectedResult()
		if actualResult != expectedResult {
			name, _ := currProblem.getInfo()
			t.Errorf("Error in problem %v (%v): solution is %v and should be %v", i+1, name, actualResult, expectedResult)
		}
	}
}
