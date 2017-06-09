package problems

import (
	"testing"
)

// Test00002 - Test Euler Problem 2
func Test00003(t *testing.T) {
	solution := "6857"

	p3 := problem00003{}
	p3.solve()

	result := p3.getResult()
	if result != solution {
		t.Errorf("Problem 3 solution is %v and should be %v", result, solution)
	}
}
