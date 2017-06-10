package problems

import "fmt"

// EulerProblem - Generic interface for an Euler Problem solution
type EulerProblem interface {
	solve()
	getInfo() (string, string)
	getExpectedResult() float64
	getActualResult() float64
}

// Problems - Array containing interface implementations of all solutions
var Problems = []EulerProblem{
	&problem00001{},
	&problem00002{},
	&problem00003{},
	&problem00004{},
	&problem00005{},
	&problem00006{},
}

// SolveAll - solve all Euler Problems and display solution
func SolveAll() {
	fmt.Printf("========================================\n")
	for i, currProblem := range Problems {
		currProblem.solve()
		name, descrip := currProblem.getInfo()
		fmt.Printf("\n")
		fmt.Printf("Problem %v - %v\n", i+1, name)
		fmt.Printf("%v\n", descrip)
		fmt.Printf("\n")
		fmt.Printf("Result: %v\n", currProblem.getActualResult())
		fmt.Printf("\n")
		fmt.Printf("========================================\n")
	}
}
