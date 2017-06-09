package problems

import (
	"fmt"
)

// EulerProblem - Generic interface for an Euler Problem solution
type EulerProblem interface {
	getName() string
	solve()
	getResult() string
}

// Problems - Array containing interface implementations of all solutions
var Problems = []EulerProblem{
	&problem00001{},
	&problem00002{},
}

// SolveAll - solve all Euler Problems and display solution
func SolveAll() {
	for i, currProblem := range Problems {
		currProblem.solve()
		fmt.Printf("Problem %d (%s): %v\n", i+1, currProblem.getName(), currProblem.getResult())
	}
}
