package problems

import (
	"testing"
)

// Test00002 - Test Euler Problem 2
func Test00002(t *testing.T) {
	solution := "4613732"

	p2 := problem00002{}
	p2.solve()

	result := p2.getResult()
	if result != solution {
		t.Errorf("Problem 2 solution is %v and should be %v", result, solution)
	}
}
