package main

import "testing"

func TestAddition(t *testing.T) {
	if calculation(1, 3, AddOperator) != 4 {
		t.Error("Expected result is 4")
	}
}

func TestSubtraction(t *testing.T) {
	if calculation(14, 17, SubOperator) != -3 {
		t.Error("Expected result is -3")
	}
}

func TestMultiplication(t *testing.T) {
	if calculation(4, 5, MulOperator) != 20 {
		t.Error("Expected result is 20")
	}
}

func TestDivision(t *testing.T) {
	if calculation(3, 4, DivOperator) != 0 {
		t.Error("Expected result is 0")
	}
}
