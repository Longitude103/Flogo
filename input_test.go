package Flogo

import "testing"

// Testing package for the input file generator.
func TestInput(t *testing.T) {
	if err := Input(true, false, true); err != nil {
		t.Error("Function produced an error")
	}

}
