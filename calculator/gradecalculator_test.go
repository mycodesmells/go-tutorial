package calculator_test

import (
	"testing"

	"github.com/mycodesmells/go-tutorial/calculator"
)

func TestFinalGradeShouldEqualToFive(t *testing.T) {
	fg := calculator.FinalGrade([]int{5})

	if(fg != 5) {
		t.Errorf("Calculated incorrect final grade. Expected 5 but was %v", fg)
	}
}
