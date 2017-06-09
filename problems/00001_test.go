package problems

import (
	"testing"
)

// Test00001 - Test Euler Problem 1
func Test00001(t *testing.T) {
	solution := "233168"

	p1 := problem00001{}
	p1.solve()

	result := p1.getResult()
	if result != solution {
		t.Errorf("Problem 1 solution is %v and should be %v", result, solution)
	}
}
