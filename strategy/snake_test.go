package strategy

import "testing"

const MAX = 3

func TestGetPrevious01(t *testing.T) {
	x, y, err := getPrevious(0, 1, MAX)
	if !(x == 1 && y == 1 && err == nil) {
		t.Error("[0, 1] didn't return properly")
	}
}

func TestGetPrevious00(t *testing.T) {
	x, y, err := getPrevious(0, 0, MAX)
	if !(x == -1 && y == -1 && err != nil) {
		t.Error("[0, 0] didn't return properly")
	}
}

func TestGetPrevious10(t *testing.T) {

	x, y, err := getPrevious(1, 0, MAX)
	if !(x == 0 && y == 0 && err == nil) {
		t.Error("[1, 0] didn't return properly")
	}
}

func TestGetPrevious30(t *testing.T) {

	x, y, err := getPrevious(3, 0, MAX)
	if !(x == 2 && y == 0 && err == nil) {
		t.Error("[3, 0] didn't return properly")
	}
}

func TestGetPrevious31(t *testing.T) {

	x, y, err := getPrevious(3, 1, MAX)
	if !(x == 3 && y == 0 && err == nil) {
		t.Error("[3, 1] didn't return properly")
	}
}

func TestGetPrevious02(t *testing.T) {

	x, y, err := getPrevious(0, 2, MAX)
	if !(x == 0 && y == 1 && err == nil) {
		t.Error("[0, 2] didn't return properly")
	}
}

func TestGetPrevious11(t *testing.T) {
	x, y, err := getPrevious(1, 1, MAX)
	if !(x == 2 && y == 1 && err == nil) {
		t.Error("[1, 1] didn't return properly")
	}
}
