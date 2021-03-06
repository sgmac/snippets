package main

import "testing"

func TestAddNumbers(t *testing.T) {
	result := AddNumbers(1, 3)
	t.Log("AddNumbers(1,3) should return 4")
	if result != 4 {
		t.Errorf("expected get 4, I got %v\n", result)
	}
}

func TestRestNumbers(t *testing.T) {
	result := RestNumbers(1, 3)
	t.Log("RestNumbers(1,3) should return -2")
	if result != -2 {
		t.Errorf("expected get -2, I got %v\n", result)
	}
}

func TestDivNumbers(t *testing.T) {
	result, err := DivNumbers(4, 2)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Log("DivNumbers(4,2) should return 2")
	if result != 2 {
		t.Errorf("expected get 2, I got %v\n", result)
	}
}

func TestDivNumbersByZero(t *testing.T) {
	_, err := DivNumbers(4, 0)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Log("DivNumbers(4,0) should return error")
}
