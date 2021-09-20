package MF6

import "testing"

// tests of the wel file generator

func TestWel(t *testing.T) {
	var data = []bool{true, false, true}

	if err := Wel("test", data); err != nil {
		t.Error("Wel function errored with", err)
	}
}
