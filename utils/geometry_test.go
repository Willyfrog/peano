package utils

import "testing"

func TestBetween(t *testing.T) {
	if !Between(0.5, 0.0, 1.0) {
		t.Error("0.5 is not between 0.0 and 1.0? Madness!")
	}
	if !Between(0.0, 0.0, 1.0) {
		t.Error("0.0 isn't between 0.0 and 1.0? Madness!")
	}
	if !Between(1.0, 0.0, 1.0) {
		t.Error("1.0 is between 0.0 and 1.0? It shouldn't be!")
	}
}
